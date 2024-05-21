package main

import "fmt"

func twoSum(nums []int, target int) []int {
	h := map[int]int{}

	for i, v := range nums {
		if t, f := h[target-v]; f {
			return []int{t, i}
		}
		h[v] = i
	}
	return []int{}
}

func main() {
	// fmt.Printf("%v\n", twoSum([]int{2, 7, 11, 15}, 9))
	// fmt.Printf("%v\n", twoSum([]int{3, 2, 4}, 6))
	fmt.Printf("%v\n", twoSum([]int{3, 3}, 6))
}
