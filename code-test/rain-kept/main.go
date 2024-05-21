package main

import (
	"fmt"
	"math/rand"
)

func main() {

	size := 5
	ls := make([]int, size)
	for i := 0; i < size; i++ {
		ls[i] = rand.Intn(100)
	}

	for i := 0; i < size; i++ {
		fmt.Printf("%d,", ls[i])
	}
	fmt.Println()

	sum := 0
	for i, j := 0, len(ls)-1; i < j; {
		m := min(ls[i], ls[j])
		for k := i; k <= j; k++ {
			ls[k] -= m
		}
		for ; ls[j] <= 0 && i < j; j-- {
			sum += ls[j]
		}
		for ; ls[i] <= 0 && i < j; i++ {
			sum += ls[i]
		}
	}

	fmt.Printf("total: %d\n", -sum)
}
