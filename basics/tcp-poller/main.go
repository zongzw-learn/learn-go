package main

import (
	"context"
	"log"
	"os"
	"syscall"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("%s <ip/domain> <port>", os.Args[0])
	}

	var np NetPoller
	var err error
	if np, err = NewKqueuePoller(os.Args[1], os.Args[2]); err != nil {
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
