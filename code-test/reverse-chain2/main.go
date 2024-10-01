package main

import (
	"fmt"
	"math/rand"
)

func main() {
	c := randChain(13)
	printChain(c)

	cc := reverseKGroup(c, 2)
	printChain(cc)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func printChain(c *ListNode) {
	for i := c; i != nil; i = i.Next {
		fmt.Printf("%d -> ", i.Val)
	}

	fmt.Println()
}

func randChain(s int) *ListNode {
	c := ListNode{Val: rand.Intn(100)}
	l := &c
	for i := 0; i < s; i++ {
		n := ListNode{Val: rand.Intn(100)}
		l.Next = &n
		l = &n
	}
	return &c
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}
	var nHead *ListNode = &ListNode{}
	var nStart, nEnd *ListNode = nHead, nHead
	nextGroup := head
	for {
		i := 0
		for ; i < k && nextGroup != nil; i++ {
			nextGroup = nextGroup.Next
		}
		if i < k {
			nEnd.Next = head
			break
		}

		nEnd = head
		for i := 0; i < k; i++ {
			t := head
			head = head.Next
			t.Next = nStart.Next
			nStart.Next = t
		}
		nStart = nEnd
	}

	return nHead.Next
}
