package main

import (
	"log"
	"time"
)

/*
	两种迭代器的实现
*/

// 采用index的方式实现(非线程安全)
type ListStruct struct {
	index int
	data  []int
}

func (sl *ListStruct) Next() int {
	d := sl.data[sl.index]
	sl.index += 1
	return d
}

func (sl *ListStruct) HasNext() bool {
	return sl.index < len(sl.data)
}

// 采用channel的方式实现
type ListType []int

func Iterator(s ListType) <-chan int {
	c := make(chan int)
	go func() {
		for _, a := range s {
			c <- a * a
		}
		close(c)
	}()
	return c
}

func main() {
	sl := ListStruct{data: []int{1, 2, 3, 4, 5}}
	for sl.HasNext() {
		log.Println(sl.Next())
	}

	s := []int{1, 2, 3, 4, 5}
	var c <-chan int

	// 使用原有数据集进行遍历（不推荐）
	c = Iterator(s)
	for i := 0; i < len(s); i++ {
		a := <-c
		log.Println(a)
		<-time.After(time.Millisecond * 10)
	}

	// 采用迭代器本身遍历（注意这里的range参数为c，而不是<-c）
	c = Iterator(s)
	for a := range c {
		log.Println(a)
		<-time.After(time.Millisecond * 10)
	}

	// 采用迭代器本身遍历（注意这里的ok判断chan结束与否）
	c = Iterator(s)
	for {
		a, ok := <-c
		if !ok {
			break
		}
		log.Println(a)
		<-time.After(time.Millisecond * 10)
	}
}
