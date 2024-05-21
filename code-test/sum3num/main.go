package main

import "fmt"

// 找出序列中加和为0的三数

func legacy(a []int) {
	for i := 0; i < len(a)-2; i++ {
		for j := i + 1; j < len(a)-1; j++ {
			for k := j + 1; k < len(a); k++ {
				if a[i]+a[j]+a[k] == 0 {
					fmt.Printf("%d %d %d\n", a[i], a[j], a[k])
				}
			}
		}
	}
}

func swap(a, b *int) {
	*b, *a = *a, *b
}

func heapOnce(a []int, n int) {
	for i := (n - 1) / 2; i >= 0; i-- {
		f := i
		s := 2*i + 1
		if i*2+2 < n && a[i*2+1] < a[i*2+2] {
			s = 2*i + 2
		}
		if s < n && a[s] > a[f] {
			swap(&a[s], &a[f])
		}
	}
}

func heap(a []int) {
	for i := len(a); i > 0; i-- {
		heapOnce(a, i)
		swap(&a[0], &a[i-1])
	}
}

func print(a []int) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d, ", a[i])
	}
	fmt.Println()
}

func sum3(a []int) {
	for i := 0; i < len(a)-2; i++ {

		if a[i] > 0 {
			return
		}
		left := i + 1
		right := len(a) - 1

		for left < right {

			if a[i]+a[left]+a[right] > 0 {
				right--
			} else if a[i]+a[left]+a[right] < 0 {
				left++
			} else {
				fmt.Printf("%d %d %d\n", a[i], a[left], a[right])
				left++
				right--
				continue
			}
			for left+1 < right && a[left] == a[left+1] {
				left++
			}
			for right-1 > left && a[right-1] == a[right] {
				right--
			}

		}
		// if i+1 < len(a) && a[i] == a[i+1] {
		// 	i++
		// 	continue
		// }
	}
}

// 解题思路：
//
//	先排序，然后循环前len-2个数字i，设置l, r
//	如果 [i]+[l]+[r] > 指定数字 r--,否则 l--,
//	其中可能会有些优化举措
//	以此，可以实现 4个数字之和为特定数的算法
func main() {
	a := []int{1, 2, -1, -2, 0, 0, 0}
	// a := []int{1, 2, 3, -1, -2, -3, 0, 0, 0}
	// a := []int{0, 0, 0}
	// a := []int{1, 0, 0, 0, -1}
	heap(a)
	print(a)

	sum3(a)
}
