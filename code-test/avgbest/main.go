package main

import (
	"fmt"
	"strings"
)

func main() {
	// var a int

	// fmt.Scan(&a)

	// r := bufio.NewReader(os.Stdin)
	// s, _ := r.ReadString('\n')
	// fmt.Printf("read: %s\n", s)
	// s = strings.Trim(s, "\n")
	// fmt.Printf("read: %s\n", s)

	// sa := strings.Split(s, " ")
	// l := make([]int, len(sa))
	// for i := 0; i < len(sa); i++ {
	// 	l[i], _ = strconv.Atoi(sa[i])
	// }

	a := 2
	// l := []int{1, 2, 3, 4, 2}
	// l := []int{1, 2, 1, 2, 5, 100, 3, 2, 3, 2, 2, 1, 1, 1}
	l := []int{1, 1, 0, 0, 2, 100, 99, 100, 2, 2, 3, 2, 0, 0, 100, 0, 0, 0, 0}
	fmt.Printf("%v\n", l)

	// 14:00
	find(l, a)
}

func below(a []int, i, j int, avg int) bool {
	s := 0
	for k := i; k <= j; k++ {
		s += a[k]
	}
	return s <= (j-i+1)*avg
}

// 2: 1 1 0 0 2 100 99 100, 2, 2, 3, 2, 0, 0, 100, 0, 0, 0, 0
func find(a []int, avg int) {
	rlt := []string{}
	for i := 0; i < len(a); i++ {
		for j := len(a) - 1; j > i; j-- {
			if below(a, i, j, avg) {
				rlt = append(rlt, fmt.Sprintf("%d-%d", i, j))
				i = j + 1
			}
		}
	}
	fmt.Printf("%s\n", strings.Join(rlt, ","))
}
