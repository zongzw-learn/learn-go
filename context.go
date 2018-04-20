package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, timeout := context.WithTimeout(context.Background(), 5*time.Second)
	defer timeout()

	ch := make(chan int)

	go func() {
		time.Sleep(10 * time.Second)
		ch <- 3
	}()

	select {
	case <-ch:
		fmt.Printf("data: finished\n")
		return
	case <-ctx.Done():
		fmt.Printf("data: timeout\n")
		return
	}

}
