package main // package treeutils

import (
	"fmt"
	"math"
)

type Node struct {
	v int
	l *Node
	r *Node
}

func height(t *Node) int {
	if t == nil {
		return 0
	}
	if t.l == nil && t.r == nil {
		return 1
	}
	return 1 + max(height(t.l), height(t.r))
}

func fmtInWidth(d []string, w int) string {
	if w <= 0 {
		return fmt.Sprintf("%v", d)
	}
	if len(d) == 0 {
		s := ""
		for i := 0; i < w; i++ {
			s += " "
		}
		return s
	}
	l := len(d)
	subw := w / l
	rlt := ""
	for _, i := range d {
		li := len(i)
		s := ""
		for i := 0; i < subw/2-li/2; i++ {
			s += " "
		}
		s += i
		for i := len(s); i < subw; i++ {
			s += " "
		}
		rlt += s
	}
	return rlt
}

func traverseInLayer(t *Node) {
	ls := []*Node{t}
	s := 0
	e := len(ls)
	h := height(t)
	width := int(math.Pow(2, float64(h-1))) * 3
	for l := 0; l < height(t); l++ {
		ds := []string{}
		for i := s; i < e; i++ {
			if ls[i] == nil {
				ds = append(ds, ".")
				ls = append(ls, nil, nil)
			} else {
				ds = append(ds, fmt.Sprintf("%d", ls[i].v))
				ls = append(ls, ls[i].l, ls[i].r)
			}
		}
		fmt.Printf("%s\n", fmtInWidth(ds, width))
		s = e
		e = len(ls)
	}
}
func main() {
	n8 := Node{v: 8}
	n7 := Node{v: 7, r: &n8}
	n6 := Node{v: 6, l: &n7}
	n5 := Node{v: 5}
	n4 := Node{v: 4, r: &n5}
	n3 := Node{v: 3, r: &n6}
	n2 := Node{v: 2, l: &n4}
	n1 := Node{v: 1, l: &n2, r: &n3}

	h := height(&n1)
	w := int(math.Pow(2, float64(h-1)))

	fmt.Printf("tree : w: %d, h: %d\n", w, h)

	// fmt.Printf("%s\n", fmtInWidth([]int{2, 3}, 8))
	// fmt.Printf("%s\n", fmtInWidth([]int{22, 33}, 8))
	// fmt.Printf("%s\n", fmtInWidth([]int{222, 333}, 8))
	// fmt.Printf("%s\n", fmtInWidth([]int{2222, 3333}, 8))
	// fmt.Printf("%s\n", fmtInWidth([]int{2, 3}, 10))
	// fmt.Printf("%s\n", fmtInWidth([]int{2, 3}, 20))

	traverseInLayer(&n1)

}
