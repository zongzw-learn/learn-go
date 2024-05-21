package main

import (
	"fmt"
)

// gene refect

func printRect(a [][]int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			fmt.Printf("%4d", a[i][j])
		}
		fmt.Println()
	}
}
func main() {
	// a := 0
	// b := 0
	// for {
	//     n, _ := fmt.Scan(&a, &b)
	//     if n == 0 {
	//         break
	//     } else {
	//         fmt.Printf("%d\n", a + b)
	//     }
	// }
	var size int
	// n, _ := fmt.Scan(&size)
	// if n == 0 {
	// 	fmt.Printf("invalid size\n")
	// 	os.Exit(1)
	// }

	// c := ""
	// n, _ = fmt.Scan(&c)
	// // fmt.Printf("%s\n", c)
	// cs := strings.Split(c, ",")
	// // fmt.Printf("%v\n", cs)
	// gr := make([]int, len(cs))
	// for i := 0; i < len(cs); i++ {
	// 	gr[i], _ = strconv.Atoi(cs[i])
	// }

	size = 5
	gr := []int{1, 2}

	a := make([][]int, size)
	for i := 0; i < size; i++ {
		a[i] = make([]int, size)
	}

	a[0] = []int{1, 1, 0, 1, 0}
	a[1] = []int{1, 1, 0, 0, 0}
	a[2] = []int{0, 0, 1, 0, 1}
	a[3] = []int{1, 0, 0, 1, 0}
	a[4] = []int{0, 0, 1, 0, 1}

	// printRect(a)

	checked := map[int]int{}
	fs := map[int]int{}
	for _, c := range gr {
		fgr := findcfor(a, c, checked)
		for _, f := range fgr {
			fs[f] = 1
		}
	}

	for _, f := range gr {
		delete(fs, f)
	}

	fmt.Println(len(fs))

}

func findcfor(a [][]int, p int, checked map[int]int) []int {
	cs := []int{}
	for i := 0; i < len(a[p]); i++ {
		if p != i && a[p][i] == 1 {
			if _, f := checked[i]; !f {
				cs = append(cs, i)
				checked[i] = 1
			}
		}
	}
	for _, c := range cs {
		subcs := findcfor(a, c, checked)
		cs = append(cs, subcs...)
	}
	return cs
}
