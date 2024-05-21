package main

import "fmt"

// 依据给定数组遍历所有房间

func main() {
	// nextVisit := []int{0, 1, 2, 0}
	// nextVisit := []int{0, 0, 2}
	nextVisit := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0} // timeout
	fmt.Printf("\n%d\n", firstDayBeenInAllRooms(nextVisit))
}

func firstDayBeenInAllRooms(nextVisit []int) int {
	m := map[int]int{}
	for i := 0; i < len(nextVisit); i++ {
		m[i] = 0
	}
	a := map[int]int{}

	for d, v := 0, 0; ; d++ {
		m[v]++
		a[v] = 1

		fmt.Printf("%d ", v)
		if len(a) == len(m) {
			return d
		}

		if m[v] > 0 && m[v]%2 == 0 {
			v = (v + 1) % len(m)
		} else {
			v = nextVisit[v]
		}
	}
}
