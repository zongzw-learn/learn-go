package main

import "fmt"

type A struct {
	a int
}

// no change
func X(a A) {
	a.a = 10
}

// change
func Y(a map[int]int) {
	a[2] = 20
}

// change
func Z(a []int) {
	a[0] = 2
}

// func D(a string) {
// 	a = "abc"
// }

func main() {
	a := A{a: 1}
	X(a)

	fmt.Println(a.a) // no change: a.a == 1

	b := map[int]int{1: 10}
	Y(b)
	fmt.Println(b) // change: b == map[int]{1:10, 2:20}

	c := []int{1}
	Z(c)
	fmt.Println(c) // change: a ==[2]

	// d := "xyz"
	// D(d)
	// fmt.Println(d) // no change: d == xyz

	s := "a"
	s += "b"
	s = "c" + s
	fmt.Println(s)
}
