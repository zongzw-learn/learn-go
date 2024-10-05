package main

import "fmt"

func main() {
	// edges := [][]int{{0, 1, 10}, {1, 2, 10}, {2, 5, 10}, {0, 3, 1}, {3, 4, 10}, {4, 5, 15}}
	// passingFees := []int{5, 1, 2, 20, 20, 3}
	maxTime := 30
	edges := [][]int{{0, 1, 10}, {1, 2, 10}, {2, 5, 10}, {0, 3, 1}, {3, 4, 10}, {4, 5, 15}}
	passingFees := []int{5, 1, 2, 20, 20, 3}

	fmt.Println(minCost(maxTime, edges, passingFees))
}

func minCost(maxTime int, edges [][]int, passingFees []int) int {
	n := len(passingFees)
	f := make([][]int, maxTime+1)
	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			f[i][j] = -1
		}
	}

	f[0][0] = passingFees[0]
	for t := 1; t <= maxTime; t++ {
		for _, edge := range edges {
			i, j, cost := edge[0], edge[1], edge[2]
			if cost <= t {
				if f[t-cost][i] != -1 {
					if f[t][j] == -1 {
						f[t][j] = f[t-cost][i] + passingFees[j]
					} else {
						f[t][j] = min(f[t][j], f[t-cost][i]+passingFees[j])
					}
				}
				if f[t-cost][j] != -1 {
					if f[t][i] == -1 {
						f[t][i] = f[t-cost][j] + passingFees[i]
					} else {
						f[t][i] = min(f[t][i], f[t-cost][j]+passingFees[i])
					}
				}
			}
		}
	}

	t := 1
	for ; t <= maxTime && f[t][n-1] == -1; t++ {
	}
	if t > maxTime {
		return -1
	} else {
		ans := f[t][n-1]
		for ; t <= maxTime; t++ {
			if f[t][n-1] != -1 && ans > f[t][n-1] {
				ans = f[t][n-1]
			}
		}
		return ans
	}
}

// func minCost(maxTime int, edges [][]int, passingFees []int) int {
// 	pathMap := cityMap(edges)
// 	passed := map[int]int{0: passingFees[0]}
// 	return travel(passed, pathMap, 0, maxTime, passingFees)
// }

// func travel(passed map[int]int, pathMap map[int]map[int]int, from int, maxTime int, passingFees []int) int {
// 	cs := []int{}
// 	cost := passed[from]

// 	for t := range pathMap[from] {
// 		if _, f := passed[t]; f {
// 			continue
// 		}
// 		if maxTime-pathMap[from][t] < 0 {
// 			continue
// 		}

// 		if t == len(passingFees)-1 {
// 			cs = append(cs, cost+passingFees[t])
// 			continue
// 		} else {
// 			passed[t] = cost + passingFees[t]
// 			leftTime := maxTime - pathMap[from][t]
// 			c2 := travel(passed, pathMap, t, leftTime, passingFees)
// 			cs = append(cs, c2)
// 			delete(passed, t)
// 		}
// 	}

// 	return minValueAbove0(cs)
// }

// // minValueAbove0 return the least value above 0, if not exist, return -1
// func minValueAbove0(a []int) int {
// 	i := 0
// 	for ; i < len(a) && a[i] < 0; i++ {
// 	}
// 	if i < len(a) {
// 		c := a[i]
// 		for j := i + 1; j < len(a); j++ {
// 			if a[j] < c {
// 				c = a[j]
// 			}
// 		}
// 		return c
// 	} else {
// 		return -1
// 	}
// }

// func cityMap(edges [][]int) map[int]map[int]int {
// 	pathMap := map[int]map[int]int{}
// 	for _, e := range edges {
// 		f, t := e[0], e[1]
// 		if _, ok := pathMap[f]; !ok {
// 			pathMap[f] = map[int]int{}
// 		}
// 		pathMap[f][t] = e[2]
// 		if _, ok := pathMap[t]; !ok {
// 			pathMap[t] = map[int]int{}
// 		}
// 		pathMap[t][f] = e[2]
// 	}

// 	return pathMap
// }
