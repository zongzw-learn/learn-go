package main

import "fmt"

func main() {
	a := []int{2, 3, 5, 1, 3, 8, 6, 4, 9, 2}
	HeapSort(a)

	fmt.Println(a)
}

func HeapSort(a []int) {
	heapOnce := func(a []int, l int) {
		f := (l - 1) / 2
		for f >= 0 {
			s1 := f*2 + 1
			s2 := f*2 + 2
			if s2 < l && a[s1] < a[s2] {
				a[s1], a[s2] = a[s2], a[s1]
			}
			if s1 < l && a[f] < a[s1] {
				a[f], a[s1] = a[s1], a[f]
			}
			f--
		}

		a[0], a[l-1] = a[l-1], a[0]
	}
	for i := len(a); i > 0; i-- {
		heapOnce(a, i)
	}
}
