package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"syscall"
	"unsafe"
)

func NewKqueuePoller(host, port string) (NetPoller, error) {
	np := KqueuePoller{}
	np.handlers = map[int]Handler{}
	np.events = make([]syscall.Kevent_t, KEVENT_NUM)
	if kq, err := syscall.Kqueue(); err != nil {
		return nil, err
	} else {
		np.kq = kq
	}
	log.Printf("kqueue fd: %d", np.kq)

	p, err := strconv.Atoi(port)
	if err != nil {
		return nil, err
	}
	addr, err := net.ResolveIPAddr("ip", host)
	if err != nil {
		return nil, err
	}
	sa := syscall.SockaddrInet4{
		Port: p,
		Addr: [4]byte(addr.IP),
	}
	// sa, err := net.ResolveTCPAddr("tcp4", "127.0.0.1")

	if fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0); err != nil {
		return nil, err
	} else {
		np.socket = fd
	}

	if err := syscall.Bind(np.socket, &sa); err != nil {
		return nil, err
	}

	if err := syscall.Listen(np.socket, BACKLOG_NUM); err != nil {
		return nil, err
	}

	if err := np.setEvent(np.socket, syscall.EVFILT_READ|syscall.EVFILT_WRITE, syscall.EV_ADD, nil); err != nil {
		return nil, err
	}

	return &np, nil
}

func (np *KqueuePoller) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			np.Close()
			return nil
		default:
			n, err := syscall.Kevent(np.kq, nil, np.events, nil)
			if err != nil {
				if err == syscall.EINTR {
					log.Printf("interrupted by external signal")
					continue
				} else {
					return err
				}
			}
			for i := 0; i < n; i++ {
				ev := np.events[i]
				if ev.Ident == uint64(np.socket) {
					var clientfd int
					var clientsa syscall.Sockaddr
					var err error
					for {
						clientfd, clientsa, err = syscall.Accept(int(ev.Ident))
						if err == nil {
							break
						} else if err == syscall.EINTR {
							log.Println("accepting " + err.Error())
						} else {
							log.Printf("failed to accept connection: %s", err.Error())
						}
					}

					// log.Printf("accepted: %d, %v\n", fdClient, clientsa)
					log.Printf("accepted connection from %v", clientsa)

					syscall.SetNonblock(clientfd, true)
					syscall.CloseOnExec(clientfd)

					// TODO: determine the SocketOps dynamically
					cs := ClientSocket{
						Fd:       clientfd,
						SockAddr: clientsa,
					}
					for filter := range np.handlers {
						if err := np.setEvent(clientfd, filter, syscall.EV_ADD, unsafe.Pointer(&cs)); err != nil {
							log.Printf("failed to set event: %s", err.Error())
						}
					}
				} else {
					// ev: syscall.Kevent_t {Ident: 7, Filter: -1, Flags: 1, Fflags: 0, Data: 7, Udata: *7}
					client := (*ClientSocket)(unsafe.Pointer(ev.Udata))
					var handler Handler
					for f, h := range np.handlers {
						if int16(f)&ev.Filter == ev.Filter {
							handler = h
						}
					}
					if handler == nil {
						log.Printf("handler for %d is not defined", ev.Filter)
					}

					switch ev.Filter {
					case syscall.EVFILT_READ:
						b, err := np.readFd(int(ev.Ident))
						if err != nil {
							if err == io.EOF {
								for filter := range np.handlers {
									if err := np.setEvent(int(ev.Ident), filter, syscall.EV_DELETE, nil); err != nil {
										log.Printf("failed to close connection: %s", err.Error())
									}
								}
								syscall.Close(int(ev.Ident))
							} else {
								log.Printf("failed to read from socket %d: %s", ev.Ident, err.Error())
							}
						} else {
							err := handler(client, b)
							if err != nil {
								log.Printf("failed to handle data from socket %d", ev.Ident)
							}
						}
					case syscall.EVFILT_WRITE:
						log.Printf("not supported for Filter type: %d", ev.Filter)
						np.setEvent(int(ev.Ident), syscall.EVFILT_WRITE, syscall.EV_DELETE, nil)
					default:
						log.Printf("not supported for Filter type: %d", ev.Filter)
					}
				}
			}
		}
	}
}

func (np *KqueuePoller) Close() {
	syscall.Close(np.socket)
	syscall.Close(np.kq)
}

func (np *KqueuePoller) SetHandler(filter int, h Handler) {
	np.handlers[filter] = h
}

func (np *KqueuePoller) readFd(fd int) ([]byte, error) {
	b := []byte{}
	for {
		buf := make([]byte, BUFF_SIZE)

		n, err := syscall.Read(fd, buf)
		if err != nil {
			if err == syscall.EAGAIN {
				break
			} else if err == syscall.EINTR {
				// 读取socket过程中碰到interrupted system call，直接跳转下一次for loop
				log.Printf("interrupted while reading")
			} else {
				return nil, err
			}
		} else if n == 0 {
			log.Printf("closing socket: %d", fd)
			return nil, io.EOF
		} else {
			b = append(b, buf[:n]...)
		}
	}
	return b, nil
}

func (np *KqueuePoller) setEvent(fd int, mode, flag int, data unsafe.Pointer) error {
	ev := syscall.Kevent_t{
		Udata: (*byte)(data),
	}
	syscall.SetKevent(&ev, fd, mode, flag)

	registered, err := syscall.Kevent(np.kq, []syscall.Kevent_t{ev}, nil, nil)
	if err != nil {
		return err
	}
	if registered == -1 {
		return fmt.Errorf("failed to register event for %d", ev.Ident)
	}
	return nil
}
