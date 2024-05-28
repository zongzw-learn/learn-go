package main

import (
	"context"
	"syscall"
)

type Handler func(*ClientSocket, []byte) error

type ClientSocket struct {
	Fd       int
	SockAddr syscall.Sockaddr
	// additional data required in future..
}

type NetPoller interface {
	Close()
	Start(ctx context.Context) error
	SetHandler(filter int, handler Handler)
}

type KqueuePoller struct {
	kq       int
	socket   int
	events   []syscall.Kevent_t
	handlers map[int]Handler
}
