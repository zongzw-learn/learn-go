package main

import "fmt"

func main() {
	days := []int{1, 4, 6, 7, 8, 20}
	// days := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}
	costs := []int{2, 7, 15}

	fmt.Println(mincostTickets(days, costs))
}

func nextXmin(days, mins []int, cur int, X int) int {
	var min, j int
	min, j = 0, cur
	nextX := days[cur] + X
	for ; j < len(days) && days[j] < nextX; j++ {
	}
	if j < len(days) {
		min = mins[j]
	}
	return min
}

func minAll(m1, m2, m3 int) int {
	m := m1
	if m > m2 {
		m = m2
	}
	if m > m3 {
		m = m3
	}
	return m
}

func mincostTickets(days []int, costs []int) int {
	mins := make([]int, len(days))

	for i := len(days) - 1; i >= 0; i-- {
		m1 := costs[0] + nextXmin(days, mins, i, 1)
		m2 := costs[1] + nextXmin(days, mins, i, 7)
		m3 := costs[2] + nextXmin(days, mins, i, 30)
		mins[i] = minAll(m1, m2, m3)
	}

	return mins[0]
}

// func mincostTickets(days []int, costs []int) int {
// 	if len(days) == 0 {
// 		return 0
// 	}
// 	m1 := costs[0] + mincostTickets(days[1:], costs)

// 	i := 0
// 	next7 := days[0] + 7
// 	for ; i < len(days) && days[i] < next7; i++ {
// 	}
// 	m2 := costs[1] + mincostTickets(days[i:], costs)

// 	j := 0
// 	next30 := days[0] + 30
// 	for ; j < len(days) && days[j] < next30; j++ {
// 	}
// 	m3 := costs[2] + mincostTickets(days[j:], costs)

// 	m := m1
// 	if m > m2 {
// 		m = m2
// 	}
// 	if m > m3 {
// 		m = m3
// 	}
// 	return m
// }
