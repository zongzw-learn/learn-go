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

func randChainReverse(s int) *Node {
	var c *Node = nil

	var n *Node
	for i := 0; i < s; i++ {
		n = &Node{data: rand.Intn(100), next: c}
		c = n
	}

	return n
}

func printChain(c *Node) {
	for i := c; i != nil; i = i.next {
		fmt.Printf("%d -> ", i.data)
	}

	fmt.Println()
}

func reverseNextBlock(current *Node, block int) (first, last, next *Node) {

	last = current
	n := current
	for i := 0; i < block && n != nil; i++ {
		c := n
		n = n.next

		c.next = first
		first = c
	}

	next = n
	return
}

func reverse(chain *Node, block int) *Node {
	//
	next := chain
	first, last, next := reverseNextBlock(next, block)
	target := first
	for next != nil {
		f, l, n := reverseNextBlock(next, block)
		last.next = f
		last = l
		next = n
	}

	return target
}

func reverseAll(chain *Node) *Node {
	var target *Node = nil

	next := chain
	for next != nil {
		c := next
		next = next.next

		c.next = target
		target = c
	}

	return target
}

// get the fix length of chain from the beginning.
func re(root *Node) (*Node, *Node) {
	c, l := root, root
	for i := 0; i < 5 && l != nil; i, l = i+1, l.next {
	}
	return c, l
}

func main() {
	// c := randChain(3)
	// printChain(c)

	d := randChainReverse(31)
	printChain(d)

	t := reverse(d, 5)
	printChain(t)

	r := reverseAll(t)
	printChain(r)

	x := randChain(20)
	c, l := re(x)
	printChain(c)
	printChain(l)
}
