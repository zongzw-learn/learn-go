package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: please provide a string")
		return
	}

	x, err := strconv.Atoi(os.Args[1])
	if err != nil || x <= 0 {
		fmt.Println("Usage: x should be an integer > 0")
		return
	}

	y := math.Sqrt(float64(x))
	_, r := math.Floor(y), math.Ceil(y)

	level := int(math.Floor(r/2 + 1))

	size := 2*(level-1) + 1
	array := make([][]int, size)
	for i := 0; i < size; i++ {
		array[i] = make([]int, size)
	}

	c := 1
	m, n := level-1, level-1
	array[m][n] = c
	c++

outer:
	for i := 2; i <= level; i++ {
		for k := 0; k < (i-1)*2; k++ {
			array[m-i+1+1+k][n+i-1] = c
			c++
			if c > x {
				break outer
			}
		}
		for k := 0; k < (i-1)*2; k++ {
			array[m+i-1][n+i-1-k-1] = c
			c++
			if c > x {
				break outer
			}
		}
		for k := 0; k < (i-1)*2; k++ {
			array[m+i-1-k-1][n-i+1] = c
			c++
			if c > x {
				break outer
			}
		}
		for k := 0; k < (i-1)*2; k++ {
			array[m-i+1][n-i+1+1+k] = c
			c++
			if c > x {
				break outer
			}
		}
	}

	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			fmt.Printf("%5d", array[i][j])
		}
		fmt.Println()
	}
}
