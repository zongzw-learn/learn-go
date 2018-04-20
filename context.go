package main

import (
	"fmt"
	"strings"
)

func heapsort(arr []int) {

	for max := len(arr) - 1; max > 0; max-- {

		for i := max / 2; i >= 0; i-- {
			ileft := 2*i + 1
			iright := 2*i + 2
			if ileft <= max && arr[ileft] > arr[i] {
				swap(arr, ileft, i)
			}
			if iright <= max && arr[iright] > arr[i] {
				swap(arr, iright, i)
			}
		}
		swap(arr, max, 0)
	}
}

func swap(arr []int, index1 int, index2 int) {
	tmp := arr[index1]
	arr[index1] = arr[index2]
	arr[index2] = tmp
}

func powered(x int) bool {
	if x == 0 {
		return false
	}
	if x&(x-1) == 0 {
		return true
	}
	return false
}

func quicksort(arr []int, s, e int) {
	m := s
	l := e
	i := s+1
	for i < l; i++ {
		for arr[i] <= arr[m] {
			i++
		}
		for arr[m] <= arr[l] {
			l--
		}
		if arr[i] > arr[m] {
			swap(arr, i, l)
		}
	}
	swap(arr, m, i)
	quicksort(arr, s, i-1)
	quicksort(arr, i+1, e)
}

func main() {
	fmt.Printf("hello world.\n")
	fmt.Printf("%s\n", strings.Join([]string{"a", "b"}, "."))
	arr := []int{5, 3, 7, 5, 3, 1, 9, 4, 2, 6, 1, 0, 16, 32, 64, 33, 65}

	fmt.Printf("%v\n", arr)
	quicksort(arr, 0, len(arr)-1)
	//heapsort(arr)
	fmt.Printf("%v\n", arr)

	for _, n := range arr {
		fmt.Printf("%d: %v\n", n, powered(n))
	}
	return
}
