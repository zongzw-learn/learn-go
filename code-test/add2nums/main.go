package main

import "fmt"

/**
 * Definition for singly-linked list.

 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	b := 0
	r := ListNode{Val: 0, Next: nil}
	l := &r
	i1, i2 := l1, l2
	for ; i1 != nil && i2 != nil; i1, i2 = i1.Next, i2.Next {
		l.Next = &ListNode{Val: (i1.Val + i2.Val + b) % 10}
		l = l.Next
		if i1.Val+i2.Val+b >= 10 {
			b = 1
		} else {
			b = 0
		}
	}
	var i *ListNode
	if i1 == nil {
		i = i2
	} else {
		i = i1
	}

	for ; i != nil; i = i.Next {
		l.Next = &ListNode{Val: (i.Val + b) % 10}
		l = l.Next
		if i.Val+b >= 10 {
			b = 1
		} else {
			b = 0
		}
	}

	if b != 0 {
		l.Next = &ListNode{Val: 1}
	}

	return r.Next
}

func fromList(a []int) *ListNode {
	r := ListNode{Val: 0}
	l := &r
	for _, i := range a {
		l.Next = &ListNode{Val: i}
		l = l.Next
	}
	return r.Next
}

func printChain(a *ListNode) {
	for l := a; l != nil; l = l.Next {
		fmt.Printf("%d -> ", l.Val)
	}
	fmt.Println()
}

func main() {
	l1 := fromList([]int{2, 4, 3})
	l2 := fromList([]int{5, 6, 4})

	printChain(l1)
	printChain(l2)

	l3 := addTwoNumbers(l1, l2)
	printChain(l3)
}
