package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Locker interface {
	Lock()
	Unlock()
}

type SpinLock int32

func (sl *SpinLock) Lock() {
	for !atomic.CompareAndSwapInt32((*int32)(sl), 0, 1) {
		// fmt.Printf("locked")
	}
}

func (sl *SpinLock) Unlock() {
	atomic.StoreInt32((*int32)(sl), 0)
}

func main() {
	l := new(SpinLock)

	v := 0
	for i := 0; i < 2; i++ {
		go func(t int) {
			for v < 1000 {
				l.Lock()
				v++
				l.Unlock()
				fmt.Printf("t%d: %d\n", t, v)
				time.Sleep(1 * time.Millisecond)
			}
		}(i)
	}

	for v < 1000 {
		<-time.After(100 * time.Millisecond)
	}
}
