package main

import (
	"fmt"
	"sync"
)

// Least Recent Used
type LRUInterface interface {
	add(a int)
	exist(a int) bool
	dump()
}

type Node struct {
	data int
	next *Node
}

type LRU struct {
	f     *Node
	size  int
	mutex sync.Mutex
}

func (x *LRU) add(a int) {
	x.mutex.Lock()
	defer x.mutex.Unlock()

	var i *Node
	for i = x.f; i != nil && i.next != nil && i.next.data != a; i = i.next {
	}

	if i == nil {
		x.f = &Node{data: a}
		return
	}
	if i.next == nil {
		c := &Node{data: a, next: x.f}
		x.f = c
		return
	}

	c := i.next
	i.next = c.next
	c.next = x.f
	x.f = c
}

func (x *LRU) exist(a int) bool {
	x.mutex.Lock()
	defer x.mutex.Unlock()

	var i *Node
	for i = x.f; i != nil && i.next != nil && i.next.data != a; i = i.next {
	}

	if i == nil || i.next == nil {
		return false
	} else {
		return true
	}
}

func (x *LRU) dump() {

	x.mutex.Lock()
	defer x.mutex.Unlock()

	for i, c := x.f, 0; i != nil && c < x.size; i = i.next {
		c++
		fmt.Printf("-> %d ", i.data)
	}
	fmt.Println()
}

func newLRU(size int) *LRU {
	return &LRU{
		mutex: sync.Mutex{},
		f:     nil,
		size:  size,
	}
}

var _ LRUInterface = (*LRU)(nil)

func main() {
	lru := newLRU(4)
	nums := []int{1, 2, 3, 4, 0, 2, 3, 4, 1, 6, 5, 2, 1, 2, 5, 4}
	for _, n := range nums {
		lru.add(n)
		lru.dump()
	}

	lrulimit()
	lruhash()
	lruhash2()
}

// limit, nums element is between 0-9
func lrulimit() {
	nums := []int{1, 2, 3, 4, 0, 2, 3, 4, 1, 6, 5, 2, 1, 2, 5, 4}
	stack := make([]int, 10)
	c := 0
	lrusize := 4
	for _, n := range nums {
		c += 1
		stack[n] = c
	}
	for i, s := range stack {
		if s > c-lrusize {
			fmt.Printf("num: %d\n", i)
		}
	}
	fmt.Println()
}

// 使用hash的方式实现保留有限的kv
func lruhash() {
	nums := []int{1, 2, 3, 4, 0, 2, 3, 4, 1, 6, 5, 2, 1, 2, 5, 4}
	s := 4

	hash := map[int]int{}
	c := 0
	l := 0
	for _, n := range nums {
		if len(hash) >= s {
			for k, v := range hash {
				if v == l {
					delete(hash, k)
					break
				}
			}
			l = (l + 1) % s
		}
		hash[n] = c
		c = (c + 1) % s
	}

	fmt.Printf("%v\n", hash)
}

// 使用hash的方式实现保留有限的kv，这是最好的方法。
func lruhash2() {
	nums := []int{1, 2, 3, 4, 0, 2, 3, 4, 1, 6, 5, 2, 1, 2, 5, 4}
	s := 4

	hash := map[int]int{}
	mark := map[int]int{}
	c := 0
	l := 0
	for _, n := range nums {
		if len(hash) >= s {
			delete(hash, mark[l])
			l = (l + 1) % s
		}
		hash[n] = c
		mark[c] = n
		c = (c + 1) % s
	}

	fmt.Printf("%v\n", mark)
}
