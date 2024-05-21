package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	data int
	next *Node
}

func randChain(s int) *Node {
	c := Node{data: rand.Intn(100)}
	l := &c
	for i := 0; i < s; i++ {
		n := Node{data: rand.Intn(100)}
		l.next = &n
		l = &n
	}
	return &c
}

func randOrderChain(s int) *Node {

	c := Node{data: rand.Intn(100)}
	l := &c
	d := c.data
	for i := 0; i < s; i++ {
		d += rand.Intn(10)
		l.next = &Node{data: d}
		l = l.next
	}

	return &c
}

func printChain(c *Node) {
	for i := c; i != nil; i = i.next {
		fmt.Printf("%d ", i.data)
		if i.next != nil {
			fmt.Printf("-> ")
		}
	}

	fmt.Println()
}

func combineChain(a, b *Node) *Node {
	c := Node{data: 0, next: nil}
	l := &c

	i, j := a, b

	for i != nil && j != nil {
		if i.data > j.data {
			l.next = j
			l = j
			j = j.next
		} else {
			l.next = i
			l = i
			i = i.next
		}
	}
	if i == nil {
		l.next = j
	}
	if j == nil {
		l.next = i
	}

	return c.next
}

func main() {
	a := randOrderChain(2)
	b := randOrderChain(3)
	printChain(a)
	printChain(b)

	c := combineChain(a, b)
	printChain(c)
}
