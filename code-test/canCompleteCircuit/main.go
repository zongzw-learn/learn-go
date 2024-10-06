package main

import "fmt"

func main() {
	// gas := []int{1, 2, 3, 4, 5}
	// cost := []int{3, 4, 5, 1, 2}

	gas := []int{2, 3, 4}
	cost := []int{3, 4, 3}

	// 0 0 0 0 1 0
	// 0 0 0 2 0 0
	fmt.Println(canCompleteCircuit(gas, cost))
}

func canCompleteCircuit(gas []int, cost []int) int {
	left := make([]int, len(gas))
	for i := 0; i < len(gas); i++ {
		left[i] = gas[i] - cost[i]
	}

	s := 0
	for j := 0; j < len(left); j++ {
		s += left[j]
	}
	if s < 0 {
		return -1
	}

	i := 0
	for {
		j := maxArrive(left, i)

		if j != i && j%len(left) == i {
			return i
		}
		i = (j + 1) % len(left)
	}

	// i := 0
	// for {
	// 	if !valid(left, i) {
	// 		i++
	// 	} else {
	// 		return i
	// 	}
	// 	if i >= len(left) {
	// 		return -1
	// 	}
	// }
}

// func valid(left []int, i int) bool {
// 	s := 0
// 	for j := i; ; {
// 		s += left[j%len(left)]
// 		if s < 0 {
// 			return false
// 		} else {
// 			j++
// 			if j%len(left) == i {
// 				return true
// 			}
// 		}
// 	}
// }

func maxArrive(left []int, i int) int {
	s := 0
	for j := i; ; {
		s += left[j%len(left)]
		if s < 0 {
			return j
		} else {
			j++
			if j%len(left) == i {
				return j
			}
		}
	}
}
