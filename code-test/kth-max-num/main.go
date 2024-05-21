package main

import (
	"fmt"
	"math/rand"
	"os"
)

func swap(a, b *int) {
	*a, *b = *b, *a
}

//         0
// 	   1 		2
// 3    4    5    6

func heapOnce(a []int, l int) {
	li := l - 1
	for i := (li - 1) / 2; i >= 0; i-- {
		f, s := i, i*2+1

		if s+1 < l && a[s] < a[s+1] {
			s = s + 1
		}
		if s < l && a[s] > a[f] {
			swap(&a[f], &a[s])
		}
	}
}
func randArray(l int) []int {
	a := make([]int, l)
	for i := 0; i < l; i++ {
		a[i] = rand.Intn(100)
	}

	return a
}

func printArray(a []int) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d  ", a[i])
	}
	fmt.Println()
}

func main() {
	a := randArray(100)
	printArray(a)
	K := 10
	if K > len(a) || K <= 0 {
		fmt.Printf("invalid K: %d\n", K)
		os.Exit(1)
	}
	l := len(a)
	for i := 0; i < K; i++ {
		heapOnce(a, l)
		swap(&a[0], &a[l-1])
		l--
	}
	printArray(a)
	fmt.Printf("Kth max: %d\n", a[len(a)-K])
}
