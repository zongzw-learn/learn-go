package main

import (
	"fmt"
	"math/rand"
)

// 计算数字序列中子串的最大累加和
func randSlice(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = 5 - rand.Intn(10)
	}

	return a
}

func printSlice(a []int) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d ", a[i])
	}
	fmt.Println()
}

func submax(a []int) int {
	temp, max := 0, 0
	for i := 0; i < len(a); i++ {
		temp += a[i]
		if temp < 0 {
			temp = 0
		}
		if temp > max {
			max = temp
		}
	}

	return max
}
func main() {
	a := randSlice(10)
	printSlice(a)

	fmt.Printf("max %d\n", submax(a))
}
