package main

import "fmt"

func main() {
	// dist := []int{1, 3, 2}
	// hours := float64(6)
	// hours := 2.7
	// hours := 1.9

	// dist := []int{1, 1, 100000}
	// hours := 2.01

	dist := []int{69}
	hours := 4.6

	fmt.Println(minSpeedOnTime(dist, hours))
}

func minSpeedOnTime(dist []int, hour float64) int {
	leastSpeed := upperInt(float64(totalL(dist)) / hour)
	maxSpeed := maxSpeed(dist, hour)

	if !validSpeed(dist, maxSpeed, hour) {
		return -1
	}

	l, r := leastSpeed, maxSpeed
	for l < r {
		mid := (l + r) / 2
		if validSpeed(dist, mid, hour) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return r
}

func validSpeed(dist []int, speed int, hour float64) bool {
	if len(dist) == 0 {
		return true
	}
	ahours := 0
	i := 0
	for ; i < len(dist)-1; i++ {
		l := dist[i]
		ahours += spent(l, float64(speed))
		if float64(ahours) > hour {
			return false
		}
	}
	return float64(ahours)+float64(dist[i])/float64(speed) <= hour
}

func spent(l int, s float64) int {
	h := float64(l) / s
	return upperInt(h)
}

func maxSpeed(dist []int, hours float64) int {
	a := 0
	for _, l := range dist {
		if l > a {
			a = l
		}
	}
	h := hours - float64(int(hours))
	b := float64(0)
	if h != float64(0) {
		b = float64(dist[len(dist)-1]) / h
	}
	if float64(a) > b {
		return a
	} else {
		return upperInt(b)
	}
}

func upperInt(a float64) int {
	x := a - float64(int(a))
	if x < 0.000000001 {
		return int(a)
	} else {
		return int(a) + 1
	}
	// if float64(int(a)) != a {
	// 	return int(a) + 1
	// } else {
	// 	return int(a)
	// }
}

func totalL(dist []int) int {
	s := 0
	for _, l := range dist {
		s += l
	}
	return s
}
