package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

type CtxKey string

var (
	KEY CtxKey = "with_value_key"
)

// 几种context的使用方法
func main() {
	// context.WithValue
	ctx1 := context.WithValue(context.Background(), KEY, "my_value")
	v := ctx1.Value(KEY)
	log.Printf("value of %s in context: %s", KEY, v)

	// context.WithTimeout
	ctx2, cancel2 := context.WithTimeout(context.Background(), time.Duration(1*time.Second))
	loopUntil(ctx2)
	defer cancel2()

	// context.WithDeadline
	ctx3, cancel3 := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	loopUntil(ctx3)
	defer cancel3()

	// context.WithCancel
	ctx4, cancel4 := context.WithCancel(context.Background())
	go loopUntil(ctx4)
	<-time.After(1 * time.Second)
	cancel4()

	//context.With*Cause
	ctx5, cancel5 := context.WithCancelCause(context.Background())
	go loopUntil(ctx5)
	cancel5(fmt.Errorf("cancelled with no reason in main thread"))
	log.Println(context.Cause(ctx5))
}

func loopUntil(ctx context.Context) {
	for {
		select {
		case <-time.After(300 * time.Millisecond):
			log.Println("running after 300 milliseconds")
		case <-ctx.Done():
			log.Println("timeout triggered")
			return
		}
	}
}
