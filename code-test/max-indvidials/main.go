package main

import "log"

func main() {
	a := []int{2, 3, 5, 4, 1, 4} // 0 2 5
	x, y := maxab(a, 0)
	log.Printf("max: %d, %v", x, y)
}

func maxa(a []int) int {
	if len(a) == 1 {
		return a[0]
	}
	if len(a) == 2 {
		return max(a[0], a[1])
	}

	s1 := a[0] + maxa(a[2:])
	s2 := maxa(a[1:])

	if s1 > s2 {
		return s1
	} else {
		return s2
	}
}

func maxab(a []int, o int) (int, []int) {
	if len(a) == 1 {
		return a[0], []int{o}
	}
	if len(a) == 2 {
		if a[0] > a[1] {
			return a[0], []int{o}
		} else {
			return a[1], []int{1 + o}
		}
	}

	x1, y1 := maxab(a[2:], 2+o)
	x2, y2 := maxab(a[1:], 1+o)
	if a[0]+x1 > x2 {
		return a[0] + x1, append([]int{o}, y1...)
	} else {
		return x2, y2
	}
}
