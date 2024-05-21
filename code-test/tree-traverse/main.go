package main

import (
	"fmt"
	"math/rand"
)

// 13:45

type Node struct {
	data        int
	left, right *Node
}

func randTree(level int) *Node {
	if level == 0 {
		return nil
	}
	r := rand.Intn(100)
	if r < 20 {
		return nil
	}

	n := Node{data: r}
	n.left = randTree(level - 1)
	n.right = randTree(level - 1)
	return &n
}

func treeDepth(t *Node, parentDepth int) int {
	c := parentDepth
	if t != nil {
		c++
		lc := treeDepth(t.left, c)
		rc := treeDepth(t.right, c)

		if lc > rc {
			return lc
		} else {
			return rc
		}
	}
	return c
}

func printNodes(ts []*Node, width int) {
	if len(ts) == 0 {
		return
	}
	sons := []*Node{}
	for _, t := range ts {
		if t != nil {
			fmt.Printf("%d  ", t.data)
			sons = append(sons, t.left, t.right)
		}
	}
	fmt.Println()
	printNodes(sons, width)

}

func main() {
	tree := randTree(5)

	d := treeDepth(tree, 0)
	fmt.Printf("tree depth: %d\n", d)

	printNodes([]*Node{tree}, d)

	// fmt.Println(tree)
}
