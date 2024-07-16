package main

import "log"

func main() {
	a := []int{2, 3, 5, 4, 1, 4} // 0 2 5
	// x, y := maxab(a, 0)
	// log.Printf("max: %d, %v", x, y)
	maxb(a)
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

// func maxb(a []int) {
// 	b := make([]int, len(a))
// 	for i := range a {
// 		if i == 0 {
// 			b[i] = a[i]
// 		} else if i == 1 {
// 			b[i] = max(a[0], a[1])
// 		} else {
// 			b1 := a[i] + b[i-2]
// 			b2 := b[i-1]
// 			b[i] = max(b1, b2)
// 		}
// 	}

// 	log.Printf("b: %v", b)
// }

func maxb(a []int) {
	// b := make([]int, len(a))
	b0, b1 := 0, 0
	b := 0
	for i := range a {
		if i == 0 {
			b0 = a[i]
			b = b0
		} else if i == 1 {
			b1 = max(a[0], a[1])
			b = b1
		} else {
			bx := b1
			b1 = a[i] + b0
			b0 = bx
			b = max(b0, b1)
		}
	}

	log.Printf("b: %v", b)
}
