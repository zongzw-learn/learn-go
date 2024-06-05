package main

import "fmt"

func main() {
	g1 := map[int][]int{
		0: {1, 3},
		1: {2},
		2: {4},
		3: {2},
		4: {0},
	}

	g2 := map[int][]int{
		0: {1, 3},
		1: {2},
		2: {4},
		3: {2},
	}

	fmt.Printf("cycle: %v\n", dag(g1))
	fmt.Printf("cycle: %v\n", dag(g2))
}

func dag(graph map[int][]int) bool {
	dots := map[int]int{}

	for k := range graph {
		if visit(dots, graph, k) {
			return true
		}
	}
	return false
}

func visit(dots map[int]int, graph map[int][]int, k int) bool {
	t := graph[k]
	if dots[k] == 1 {
		return false
	}

	if dots[k] == -1 {
		return true
	}

	dots[k] = -1
	for _, v := range t {
		if c := visit(dots, graph, v); c {
			return true
		}
	}
	dots[k] = 1

	return false
}
