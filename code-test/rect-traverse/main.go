package main

import (
	"fmt"
	"math/rand"
)

func printRect(a [][]int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			fmt.Printf("%5d", a[i][j])
		}
		fmt.Println()
	}
}

func randRect(i, j int) [][]int {
	a := make([][]int, i)
	for n := 0; n < i; n++ {
		a[n] = make([]int, j)
		for k := 0; k < j; k++ {
			a[n][k] = rand.Intn(100)
		}
	}
	return a
}

func traverse(a [][]int) {
	x, y := len(a), len(a[0])
	reverse := 0
	for k := 0; k < x+y; k++ {
		for i := 0; i < x+y; i++ {
			j := k - i
			if 0 <= i && i < y && 0 <= j && j < x {
				// if reverse == 0 {
				// fmt.Printf("(%d, %d): %d\n", i, j, a[i][j])
				// } else {
				fmt.Printf("(%d, %d): %d\n", j, i, a[j][i])
				// }
			}
		}
		reverse ^= 1
	}
}

func main() {
	a := randRect(4, 4)
	printRect(a)

	fmt.Println()
	traverse(a)
}
