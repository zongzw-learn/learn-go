package main

import (
	"context"
	"log"
	"syscall"
)

func main() {
	var np NetPoller
	var err error
	if np, err = NewKqueuePoller("0.0.0.0", "8082"); err != nil {
		log.Fatal(err)
	}

	np.SetHandler(syscall.EVFILT_READ, handleRead2)
	// np.SetHandler(syscall.EVFILT_WRITE, handleRead2)
	defer np.Close()

	ctx, cancel := context.WithCancel(context.TODO())
	err = np.Start(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer cancel()
}
