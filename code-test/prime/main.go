package main

import (
	"fmt"
	"runtime"
)

// func is_prime(exists []int, a int) bool {
// 	for _, e := range exists {
// 		if a < 2 || a%e == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func main() {

// 	ps := []int{2}

// 	cs := make(chan int, runtime.NumCPU())
// 	max := 100
// 	go func() {
// 		for i := 3; i < max; i++ {
// 			cs <- i
// 		}
// 	}()

// 	for i := 3; i < max; i++ {
// 		a := <-cs
// 		// go func() {
// 		if is_prime(ps, a) {
// 			ps = append(ps, a)
// 		}
// 		// }()
// 	}

// 	fmt.Printf("%v\n", ps)

// }

func main() {

	ps := []int{2}

	for i := 3; i < 10000; i++ {
		is := true
		cs := make(chan bool, runtime.NumCPU())
		for _, p := range ps {
			go func(p int, a int) {
				// fmt.Printf("a: %d, p: %d\n", a, p)
				cs <- a%p != 0
			}(p, i)
		}
		for p := 0; p < len(ps); p++ {
			is = is && <-cs
		}

		if is {
			// fmt.Printf("append: %d\n", i)
			ps = append(ps, i)
		} else {
			// fmt.Printf("skipped: %d\n", i)
		}
	}

	fmt.Printf("%v\n", ps)

}
