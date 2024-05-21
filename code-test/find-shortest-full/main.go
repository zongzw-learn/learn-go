package main

import "fmt"

// 找出不重复的最短子串

func found(a int, ls []int) bool {

	for _, i := range ls {
		if a == i {
			return true
		}
	}
	return false
}

func main() {
	ls := []int{1, 2, 3, 4, 2, 3, 4, 5, 1, 2, 6, 2, 4, 1, 1, 7}
	// ls := []int{1, 2, 3, 4, 2, 3, 4, 5, 1, 2}
	// ls := []int{1, 2, 3, 4, 2, 3, 4, 5, 1, 2}
	// ls := []int{1, 2, 3, 4, 2, 3, 4, 5, 1, 2}

	min := 0

	for i, j := 0, 1; j < len(ls); {
		if ls[i] == ls[j] || found(ls[i], ls[i+1:j]) {
			i++
			if min > j-i {
				min = j - i
			}
		} else {
			j++
			if j < len(ls) && !found(ls[j], ls[i:j]) {
				min = j - i
			}
		}
	}

	fmt.Printf("min length: %d\n", min+1)
}
