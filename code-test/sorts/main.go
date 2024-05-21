package main

import (
	"fmt"
	"math/rand"
)

// var swapcount = 0

func swap(a, b *int) {
	*b, *a = *a, *b
	// swapcount++
}

func randArray(len int) []int {
	a := make([]int, len)
	for i := 0; i < len; i++ {
		a[i] = rand.Intn(len * 10)
	}
	return a
}

func heapOnce(a []int, n int) {
	for i := (n - 1) / 2; i >= 0; i-- {
		f := i
		s := 2*i + 1
		if i*2+2 < n && a[i*2+1] < a[i*2+2] {
			s = 2*i + 2
		}
		if s < n && a[s] > a[f] {
			swap(&a[s], &a[f])
		}
	}
}

func heap(a []int) {
	for i := len(a); i > 0; i-- {
		heapOnce(a, i)
		swap(&a[0], &a[i-1])
	}
}

func bubble(a []int) {
	for i := len(a); i > 0; i-- {
		swapped := false
		for j := 0; j < i-1; j++ {
			if a[j] > a[j+1] {
				swapped = true
				swap(&a[j], &a[j+1])
			}
		}
		if !swapped {
			break
		}
	}
}

func quick(a []int, i, j int) {
	if i >= j {
		return
	}

	l, r := i+1, j
	for l < r {
		for ; l < r && a[l] <= a[i]; l++ {
		}
		for ; l < r && a[r] >= a[i]; r-- {
		}
		a[l], a[r] = a[r], a[l]
	}
	if a[l] < a[i] {
		a[l], a[i] = a[i], a[l]
	}
	quick(a, i, l-1)
	quick(a, l, j)
}

func print(a []int) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d, ", a[i])
	}
	fmt.Println()
}

func main() {

	// a := []int{66, 27, 71, 9, 96, 16, 64, 56, 41, 22}
	// a := []int{30, 69, 33, 81, 79, 30, 34, 69, 69, 0}
	a := randArray(10)
	quick(a, 0, len(a)-1)
	print(a)

	b := randArray(10)
	heap(b)
	print(b)

	c := randArray(10)
	bubble(c)
	print(c)
}
