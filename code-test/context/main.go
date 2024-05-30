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
	// WithValue可以使得我们在context对象中携带某种上下文数据，例如requestID。
	// 这种上下文数据可以是多个，通过多次调用WithContext实现，例如：
	// ctx := context.WithValue(context.Background(), Key1, Value1)
	// ctx = context.WithValue(ctx, Key2, Value2)
	// v1 := ctx.Value(Key1) -> Value1
	// v2 := ctx.Value(Key2) -> Value2
	// 需要注意的是Key 应该是内置类型，比如以上定义的CtxKey 类型
	ctx1 := context.WithValue(context.Background(), KEY, "my_value")
	v := ctx1.Value(KEY)
	log.Printf("value of %s in context: %s", KEY, v)

	// context.WithTimeout
	// 设置程序的执行超时时长。
	ctx2, cancel2 := context.WithTimeout(context.Background(), time.Duration(1*time.Second))
	loopUntil(ctx2)
	defer cancel2()

	// context.WithDeadline
	// WithDeadline 和WithTimeout类似，不同的是WithDeadline传入的是一个时间点，而不是时长。
	ctx3, cancel3 := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	loopUntil(ctx3)
	defer cancel3()

	// context.WithCancel
	// WithCancel用于 主动cancel控制
	ctx4, cancel4 := context.WithCancel(context.Background())
	go loopUntil(ctx4)
	<-time.After(1 * time.Second)
	cancel4()

	//context.With*Cause
	// With*Cause包括 WithCancelCause WithTimeoutCause WithDeadlineCause
	// 用于告知cancel的原因
	ctx5, cancel5 := context.WithCancelCause(context.Background())
	go loopUntil(ctx5)
	<-time.After(1 * time.Second)
	cancel5(fmt.Errorf("cancelled with no reason in main thread"))

	// delay quitting main func
	<-time.After(100 * time.Millisecond)
}

func loopUntil(ctx context.Context) {
	for {
		select {
		case <-time.After(300 * time.Millisecond):
			log.Println("running after 300 milliseconds")
		case <-ctx.Done():
			log.Printf("stop triggered, cause: %s", context.Cause(ctx))
			return
		}
	}
}
