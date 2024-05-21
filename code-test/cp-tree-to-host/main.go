package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
笔试试题：

// 定义二叉树节点结构

struct TreeNode {

    int val;

    TreeNode* left;

    TreeNode* right;



    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}

};



在机器A上有一颗树 treeA

需要把这颗树复制到机器B上

不考虑远程通信，
*/

type Node struct {
	v int
	l *Node
	r *Node
}

type LNode struct {
	v int
	l int
	r int
}

// var lnodes []LNode = []LNode{}

// func traverse(i int) {
// 	total := len(lnodes)
// 	for ; i < total; i++ {
// 		n := lnodes[i].n
// 		if n.l != nil {
// 			lnodes[i].l = len(lnodes)
// 			lnodes = append(lnodes, LNode{n: n.l})
// 		}
// 		if n.r != nil {
// 			lnodes[i].r = len(lnodes)
// 			lnodes = append(lnodes, LNode{n: n.r})
// 		}
// 	}
// 	if total < len(lnodes) {
// 		traverse(total)
// 	}
// }

// func encode(n *Node) []string {
// 	lnodes = append(lnodes, LNode{n: n})
// 	traverse(0)
// 	rlt := []string{}
// 	for _, n := range lnodes {
// 		fmt.Printf("%d(%d, %d) -> ", n.n.v, n.l, n.r)
// 		rlt = append(rlt, fmt.Sprintf("%d|%d,%d", n.n.v, n.l, n.r))
// 	}
// 	return rlt
// }

//	func newNode(s string) (*Node, int, int) {
//		v := strings.Split(s, "|")[0]
//		lr := strings.Split(s, "|")[1]
//		vv, _ := strconv.Atoi(v)
//		l, _ := strconv.Atoi(strings.Split(lr, ",")[0])
//		r, _ := strconv.Atoi(strings.Split(lr, ",")[1])
//		return &Node{v: vv}, l, r
//	}

// func decode(nodes []string, i int) *Node {
// 	root, l, r := newNode(nodes[i])
// 	if l != 0 {
// 		root.l = decode(nodes, l)
// 	}
// 	if r != 0 {
// 		root.r = decode(nodes, r)
// 	}
// 	return root
// }

func traverse(t *Node) []LNode {
	if t == nil {
		return nil
	}
	rlt := []LNode{{v: t.v}}
	if t.l != nil {
		rlt[0].l = len(rlt)
		lrlt := traverse(t.l)
		rlt = append(rlt, lrlt...)
	}
	if t.r != nil {
		rlt[0].r = len(rlt)
		rrlt := traverse(t.r)
		rlt = append(rlt, rrlt...)
	}

	return rlt
}

func enc(t *Node) []string {
	l := traverse(t)
	rlt := []string{}
	for _, i := range l {
		rlt = append(rlt, fmt.Sprintf("%d,%d,%d", i.v, i.l, i.r))
	}
	return rlt
}

func dec(s []string) *Node {
	if len(s) == 0 {
		return nil
	}
	root, l, r := mknode(s[0])
	if l != 0 {
		root.l = dec(s[l:])
	}
	if r != 0 {
		root.r = dec(s[r:])
	}
	return root
}

func mknode(s string) (*Node, int, int) {
	vv := strings.Split(s, ",")
	v, _ := strconv.Atoi(vv[0])
	l, _ := strconv.Atoi(vv[1])
	r, _ := strconv.Atoi(vv[2])
	return &Node{v: v}, l, r
}

// func main() {
// 	n6 := Node{v: 6}
// 	n5 := Node{v: 5}
// 	n4 := Node{v: 4, r: &n5}
// 	n3 := Node{v: 3, r: &n6}
// 	n2 := Node{v: 2, l: &n4}
// 	n1 := Node{v: 1, l: &n2, r: &n3}

// 	// e := encode(&n1)
// 	// fmt.Printf("%v\n", e)
// 	// d := decode(e, 0)
// 	// fmt.Printf("%v\n", d)
// 	e := enc(&n1)
// 	fmt.Printf("%v\n", e)
// 	d := dec(e)
// 	fmt.Printf("%v\n", d)
// }

type NodeState struct {
	v    int
	l, r int
}

func backup(root *Node) []NodeState {
	nss := []NodeState{}
	if root != nil {
		nss = append(nss, NodeState{v: root.v})

		lss := backup(root.l)
		if len(lss) > 0 {
			nss[0].l = 1
			nss = append(nss, lss...)
		}

		rss := backup(root.r)
		if len(rss) > 0 {
			nss[0].r = len(nss)
			nss = append(nss, rss...)
		}
	}
	return nss
}

func restore(nss []NodeState) *Node {
	if len(nss) == 0 {
		return nil
	}
	root := Node{v: nss[0].v}
	lidx, ridx := nss[0].l, nss[0].r
	if lidx > 0 {
		root.l = restore(nss[lidx:])
	}
	if ridx > 0 {
		root.r = restore(nss[ridx:])
	}
	return &root
}

func main() {
	n6 := Node{v: 6}
	n5 := Node{v: 5}
	n4 := Node{v: 4, r: &n5}
	n3 := Node{v: 3, r: &n6}
	n2 := Node{v: 2, l: &n4}
	n1 := Node{v: 1, l: &n2, r: &n3}

	b := backup(&n1)
	fmt.Printf("%v\n", b)
	r := restore(b)
	fmt.Printf("%v\n", r)
}
