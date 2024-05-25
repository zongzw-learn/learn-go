package main

import (
	"fmt"
	"time"
)

func main() {
	x := [...]int{1, 2, 3} // 数组的声明中可以使用...来代替具体的数字
	y := x                 // 数组的赋值是全拷贝方式，所以对新数组的修改不影响原始数组
	y[0] = 4
	// z := append(x, 2)   // 数组不可以使用append 操作。
	fmt.Printf("&x: %p, &y: %p\n", &x, &y)
	fmt.Printf("x: %v, y: %v\n", x, y)

	m := []int{1, 2, 3}
	n := m // 切片的赋值仅仅将索引部分拷贝，所以对新切片的修改会影响到原始切片的数据
	n[0] = 4
	fmt.Printf("&m: %p, &n: %p\n", &m, &n)
	fmt.Printf("m: %v, n: %v\n", m, n)

	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("slide len: %d, cap: %d\n", len(slice), cap(slice))          // 原始切片的长度和容量：slide len: 6, cap: 6
	subSlice := slice[1:4]                                                  // 子切片的长度和容量，注意这里的容量为新切片在原始切片位置开始的剩余容量。
	fmt.Printf("subSlide len: %d, cap: %d\n", len(subSlice), cap(subSlice)) // subSlide len: 3, cap: 5

	slice = append(slice, 7)                                       // 容量总是 以x2的方式增长。
	fmt.Printf("slide len: %d, cap: %d\n", len(slice), cap(slice)) // slide len: 7, cap: 12
	slice = append(slice, 8, 9, 10, 11, 12, 13)
	fmt.Printf("slide len: %d, cap: %d\n", len(slice), cap(slice)) // 再次翻倍： cap:24

	i := []int{1, 2, 3}
	j := i[:]
	j[0] = 4                               // 切片是对现有数据（内存块）的索引，所以，对切片内的元素的修改会影响到所以该数据的所有切片
	fmt.Printf("&i: %p, &j: %p\n", &i, &j) //
	fmt.Printf("i: %v, j: %v\n", i, j)     // i: [4 2 3], j: [4 2 3]

	// 这可能是最难理解的地方，因为cap导致的底层数据被拷贝，上层行为产生差异。
	// 如果底层数组的容量不足以容纳附加的元素，则 append() 函数将分配一个新的、更大容量的数组保存结果。
	// 如果 append 操作之后创建了一个更大的数组，则新切片将不在与原来的子切片共享相同的底层数组。
	a := []int{1, 2, 3}     // a 的长度为3， 容量为3
	b := append(a, 4, 5, 6) // b 的长度为6， 容量为6，因为原有内存位置的容量不足以保存新的容量，所以a和b元素的底层物理位置并不相同
	b[0] = 4                // 对b的修改不影响a的值
	fmt.Printf("a: %v, b: %v\n", a, b)

	c := append(a)
	c[0] = 2
	d := a[:2]
	d = append(d, 7)
	fmt.Printf("a: %v, c: %v, d: %v\n", a, c, d) // c 和d 操作均可以在原有容量中完成，所以对各个切片元素的修改会影响到a切片

	fmt.Printf("[]int{} len: %d, cap: %d\n", len([]int{}), cap([]int{})) // 空的slice的长度和容量都是0

	testCopy()
	benchmarkArrayVSSlice()
}

func testCopy() {
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 2, 4)
	n := copy(b, a)                           // copy 函数将元素复制，而不是让b索引原slice
	fmt.Printf("copied from a to b: %d\n", n) // 复制数量和原slice及目标slice的长度有关系
	fmt.Printf("b: %v\n", b)
	b[0] = 4 // 对复制后的slice改动不会影响原有slice，因为是深度拷贝
	fmt.Printf("a: %v, b: %v\n", a, b)
}

/*
重要：
	不同slice，可能会共享相同的底层位置，
	为了能够将GC将不必要的内存回收，尤其是大slice占据的内容，
	我们可以使用copy的方式将大slice中关注的部分深度拷贝存储，而不是引用大slice。
	这样，大slice就可以及时的被GC回收掉。
*/

// 值分配、参数传递、使用 range 关键字循环等，都涉及值拷贝。值大小越大，拷贝的代价就越大
// 所以，在出现性能问题的时候要考虑下是否可以使用 slice代替数组提升效率。
func benchmarkArrayVSSlice() {
	a := [10000]int{}
	b := make([]int, 10000)
	var s, e time.Time

	s = time.Now()
	for range a {
	}
	e = time.Now()
	fmt.Printf("loop array: %d\n", e.Sub(s).Nanoseconds())

	s = time.Now()
	for range b {
	}
	e = time.Now()
	fmt.Printf("loop slice: %d\n", e.Sub(s).Nanoseconds())
}
