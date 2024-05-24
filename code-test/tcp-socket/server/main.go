package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"runtime"
	"time"
)

func main() {
	l, err := net.Listen("tcp", ":"+os.Args[1])
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	go func() {
		num := runtime.NumGoroutine()
		for {
			<-time.After(2 * time.Second)
			if num != runtime.NumGoroutine() {
				num = runtime.NumGoroutine()
				fmt.Printf("current goroutine: %d\n", num)
			}
		}
	}()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("%s", err.Error())
		}
		fmt.Printf("connected from %s\n", conn.RemoteAddr().String())

		go serv(conn)
	}
}

func serv(conn net.Conn) {
	defer conn.Close()

	buff := make([]byte, 128)
	for {
		n, err := conn.Read(buff)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Printf("readed: %d\n", n)
		fmt.Println(buff[0:n])
	}
}
