package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 使用3个goroutine交替打印ABC，并且在输出是也按照 ABCABC 顺序输出。

var stopped = make(chan bool)

func test(f, t chan int) {
outer:
	for {
		select {
		case <-stopped:
			fmt.Printf("stopped.\n")
			break outer
		case i := <-f:
			fmt.Printf("%c", rune(i+int('A')))
			<-time.After(1 * time.Second)
			i = (i + 1) % 3
			t <- i
		}
	}
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println("Received signal:", sig)
		close(stopped)
	}()

	cs := make([]chan int, 4)
	for i := range cs {
		cs[i] = make(chan int)
	}
	for i := 0; i < len(cs); i++ {
		go test(cs[i], cs[(i+1)%len(cs)])
	}
	cs[0] <- 0

	<-stopped
	// <-time.After(1 * time.Second)
}
