package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"strings"
	"syscall"
)

type ClientSocket struct {
	Fd       int
	SockAddr syscall.Sockaddr
	// additional data required in future..
}

func handleRead1(cs *ClientSocket, data []byte) error {
	log.Printf("read from client: %d, %s\n", len(data), strings.Trim(string(data), "\n"))
	return nil
}

// calculate md5 and return
func handleRead2(cs *ClientSocket, data []byte) error {
	w := md5.New()
	w.Write(data)
	hex := fmt.Sprintf("MD5: %x --- %s\n", hex.EncodeToString(w.Sum(nil)), data)
	n, err := syscall.Write(cs.Fd, []byte(hex))
	if err != nil {
		return err
	}
	log.Printf("write back to client with %d bytes", n)
	return nil
}
