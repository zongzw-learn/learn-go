package main

import (
	"fmt"
)

func main() {
	inhand := []int{3, 3, 3, 4, 4, 5, 5, 6, 7, 8, 9, 10, 11, 12, 12, 12, 12, 13, 14}
	ondesk := []int{4, 5, 6, 7, 8, 8, 8}
	behind := map[int]int{}
	for i := 3; i < 15; i++ {
		behind[i] = 4
	}
	for _, i := range append(inhand, ondesk...) {
		behind[i] -= 1
	}

	fmt.Printf("%v\n", behind)

	chain := []int{}
	c := 0
	for i := 14; i > 2; i-- {
		if behind[i] > 0 {
			chain = append([]int{i}, chain...)
			c++
		} else {
			if c >= 5 {
				break
			} else {
				c = 0
				chain = []int{}
			}
		}
	}
	if len(chain) > 0 {
		printChain(chain)
	} else {
		fmt.Printf("NO CHAIN\n")
	}
}

func printChain(a []int) {
	for i := 0; i < len(a)-1; i++ {
		fmt.Printf("%d-", a[i])
	}
	fmt.Printf("%d\n", a[len(a)-1])
}
