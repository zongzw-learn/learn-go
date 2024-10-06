package main

import "fmt"

func main() {
	// a := []int{2, 3, 57, 3, 2, 1, 6, 4, 3}
	// heapSort(a)
	// fmt.Println(a)
	time := []int{1, 2, 3}
	totalTrips := 5
	fmt.Println(minimumTime(time, totalTrips))
}

func minimumTime(time []int, totalTrips int) int64 {
	maxTime := 0
	for i := 0; i < len(time); i++ {
		if time[i] > maxTime {
			maxTime = time[i]
		}
	}
	i, j := 0, maxTime*totalTrips
	for i < j {
		m := (i + j) / 2
		if valid(time, m, totalTrips) {
			j = m
		} else {
			i = m + 1
		}
	}
	return int64(i)
}

func valid(time []int, given int, trips int) bool {
	c := 0
	for _, t := range time {
		c += given / t
		if c >= trips {
			return true
		}
	}
	return false
}

// func minimumTime(time []int, totalTrips int) int64 {

// 	heapSort(time)
// 	l := len(time)
// 	t := 1
// 	for ; ; t++ {
// 		tt := t * time[l-1]
// 		c := 0
// 		for i := 0; i < len(time); i++ {
// 			if tt >= time[i] {
// 				c += tt / time[i]
// 			} else {
// 				break
// 			}
// 		}
// 		if c >= totalTrips {
// 			break
// 		}
// 	}

// 	for tt := (t - 1) * time[l-1]; tt <= t*time[l-1]; t++ {
// 		c := 0
// 		for i := 0; i < len(time); i++ {
// 			if t >= time[i] {
// 				c += t / time[i]
// 			} else {
// 				break
// 			}
// 		}
// 		if c >= totalTrips {
// 			return int64(t)
// 		}
// 	}
// 	return -1
// }

// func heapSort(a []int) {
// 	heapOnce := func(a []int, l int) {
// 		if l < 2 {
// 			return
// 		}
// 		for f := (l - 2) / 2; f >= 0; f-- {
// 			s1 := 2*f + 1
// 			s2 := 2*f + 2
// 			if s2 < l && a[s1] < a[s2] {
// 				s1 = s2
// 			}
// 			if s1 < l && a[s1] > a[f] {
// 				a[f], a[s1] = a[s1], a[f]
// 			}
// 		}

// 		a[0], a[l-1] = a[l-1], a[0]
// 	}
// 	for l := len(a); l > 0; l-- {
// 		heapOnce(a, l)
// 	}
// }
