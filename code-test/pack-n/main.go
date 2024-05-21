package main

import (
	"fmt"
	"io"
	"unsafe"
)

// https://zhuanlan.zhihu.com/p/53413177

/**
 * 结构体的成员变量，第一个成员变量的偏移量为 0。往后的每个成员变量的对齐值必须为编译器默认对齐长度（#pragma pack(n)）或当前成员变量类型的长度（unsafe.Sizeof），取最小值作为当前类型的对齐值。其偏移量必须为对齐值的整数倍
 * 结构体本身，对齐值必须为编译器默认对齐长度（#pragma pack(n)）或结构体的所有成员变量类型中的最大长度，取最大数的最小整数倍作为对齐值
 * 结合以上两点，可得知若编译器默认对齐长度（#pragma pack(n)）超过结构体内成员变量的类型最大长度时，默认对齐长度是没有任何意义的
 **/

/** 自己的理解：**/

/*
* 结构中的成员变量在内存中的排布为什么会出现对齐/填充现象：
 *	是为了让CPU在一个时钟周期内尽可能全地读取所要读取的内容，不对齐的情况下，一个数据可能会分两次读取，尤其是像int64这种占空间比较大的数据。
*/

/*
 * 什么是对齐值：
  * 不同的类型的对齐值各不相同：
    alignof bool            1
	alignof string          8
	alignof int             8
	alignof rune            4
	alignof float           8
	alignof int32           4
	alignof float32         4
	alignof []int           8
	alignof map[s]s         8
	alignof struct A        2
	alignof struct D        8
	alignof struct L        1

	在不同平台上的编译器都有自己默认的 “对齐系数”，可通过预编译命令 #pragma pack(n) 进行变更，n 就是代指 “对齐系数”。一般来讲，我们常用的平台的系数如下：

    32 位：4
    64 位：8

	另外要注意，不同硬件平台占用的大小和对齐值都可能是不一样的。因此本文的值不是唯一的，调试的时候需按本机的实际情况考虑

*/

/*
* 偏移：偏移（Offset）必须是对齐值(Align)的整数倍，对齐值的取值：系统的默认对齐值和当前成员对齐值，取小
在64位机器上，系统默认对齐值为8，一般类型的对齐值都<=此值
之所以这样做是为了让cpu能够准确从一个“整数”偏移处读取内容。
*/

/*
* 整个结构的大小必须是 对齐尺寸（取[成员最大对齐值]和[默认对齐值]中的较大值）的最小整数倍，往往都是成员最大值
**/

type A struct {
	a bool  // 偏0 占1 补1
	b int16 // 偏2 占2 补0
} // 共4个字节

type B struct {
	a bool  // 偏0 占1 补1
	b int16 // 偏2 占2 补4
	c int64 // 偏8 占8 补0 占八个字节，共12字节，但是结构需要“整除较大整数”（较大整数= 32位架构下的4 和 int64的8，取8）
} // 共16字节

type C struct {
	a bool   // 偏0 占1 补1
	b int16  // 偏2 占2 补4
	c int64  // 偏8 占8 补0
	d string // 偏16 占16 补0
} // 共32字节

type D struct {
	a bool   // 偏0 占1 补1
	b int16  // 偏2 占2 补4
	c int64  // 偏8 占8 补0
	d string // 偏16 占16 补0
	e []int  // 偏32 占24 补0  slice map 等复杂数据结构的 占位均为24，align 均为8
} // 共56字节

type E struct {
	a bool  // 偏0 占1 补7
	b []int // 偏8 占24 补0
	c D     //偏32 占56 补0
} // 共88字节

type F struct {
	a bool  // 偏0 占1 补7
	b int64 // 偏8 占8 补0
} // 16

type G struct {
	a bool // 没做任何补齐
} // 1

type H struct {
	a bool   // 偏0 占1 补7
	b []int  // 偏8 占24 补0
	c D      //偏32 占56 补0
	d string //偏88 占16 补0
} // 共104字节

type I struct {
	a bool  // 偏0 占1 补0
	b byte  // 偏1 占1 补6
	c []int // 偏8 占24 补0
} // 32字节

type J struct {
	a bool // 偏0 占1 补0
	b byte //  偏1 占1 补0
} // 2

type K struct {
	a bool // 偏0 占1 补0
	b byte //  偏1 占1 补2
	c rune // 偏4 占4 补0
} // 8

type L struct {
	a bool // 偏0 占1 补0
	b byte //  偏1 占1 补0
	c bool // 偏2 占1 补0
} // 3

type M struct {
	a bool  // 偏0 占1 补0
	b byte  //  偏1 占1 补0
	c bool  // 偏2 占1 补1
	d int16 // 偏4 占2 补0
} // 6

type N struct {
	a int16 // 偏0 占2 补0
	b L     // 偏2 占3 补1    对齐方式采用int16的整数倍补齐
} // 6

type O struct {
	a int16 // 偏0 占2 补0
	b L     // 偏2 占3 补0
	c bool  // 偏5 占1 补0
} // 6

type P struct {
	a int16 // 偏0 占2 补0
	b L     // 偏2 占3 补1
	c int16 // 偏6 占2 补0
} // 8

type Q struct {
	a int64 // 偏0 占8 补0
	b int16 // 偏8 占2 补6
} // 16

type R struct {
	a []int // 偏0 占24 补0
	b bool  // 偏24 占1 补7   补系统默认对齐值
} // 32

type S struct {
	a int16 // 偏0 占2 补0
	b L     // 偏2 占3 补0
	c bool  // 偏5 占1 补0
	d bool  // 偏6 占1 补1
} // 8

type T struct {
	a int16 // 偏0 占2 补0
	b bool  // 偏2 占1 补1  末尾的补1 是为了凑齐成员变量中较大补齐值的最小整数倍
} // 4

type U struct {
	a B    // 偏0 占16 补0
	b bool // 偏16 占1 补7  末尾的补7 是为了凑齐成员变量B的补齐值8的最小整数倍
} // 24

type V struct {
	r    io.Reader      // 偏0 占16 补0
	curr numBytesReader // 偏16 占16 补0
	skip int64          // 偏32 占8 补0
	blk  block          // 偏40 占64 补0
	err  error          // 偏104 占16 补0
} // 120
type block [64]byte
type numBytesReader [16]byte

type Part1 struct {
	a bool  // 偏0 占1 补3
	b int32 // 偏4 占4 补0
	c int8  // 偏8 占1 补7
	d int64 // 偏16 占8 补0
	e byte  // 偏24 占1 补7
} // 32 字节

type Part2 struct {
	e byte  // 偏0 占1 补0
	c int8  // 偏1 占1 补0
	a bool  // 偏2 占1 补1
	b int32 // 偏4 占4 补0
	d int64 // 偏8 占8 补0
} // 16

func main() {
	fmt.Printf("alignof bool 		%d\n", unsafe.Alignof(bool(true)))
	fmt.Printf("alignof string 		%d\n", unsafe.Alignof(string("a")))
	fmt.Printf("alignof int 		%d\n", unsafe.Alignof(int(0)))
	fmt.Printf("alignof rune 		%d\n", unsafe.Alignof(rune(0)))
	fmt.Printf("alignof float 		%d\n", unsafe.Alignof(float64(0)))
	fmt.Printf("alignof int32 		%d\n", unsafe.Alignof(int32(0)))
	fmt.Printf("alignof float32 	%d\n", unsafe.Alignof(float32(0)))
	fmt.Printf("alignof []int 		%d\n", unsafe.Alignof([]int{}))
	fmt.Printf("alignof map[s]s 	%d\n", unsafe.Alignof(map[string]string{}))
	fmt.Printf("alignof struct A	%d\n", unsafe.Alignof(A{}))
	fmt.Printf("alignof struct D	%d\n", unsafe.Alignof(D{}))
	fmt.Printf("alignof struct L	%d\n", unsafe.Alignof(L{}))
	fmt.Println()

	fmt.Printf("size of bool: 		%d\n", unsafe.Sizeof(true))
	fmt.Printf("size of rune: 		%d\n", unsafe.Sizeof('a'))

	fmt.Printf("size of string: 	%d\n", unsafe.Sizeof("AB"))
	fmt.Printf("size of string: 	%d\n", unsafe.Sizeof("ABC"))
	fmt.Printf("size of string: 	%d\n", unsafe.Sizeof("ABCD"))
	fmt.Printf("size of string: 	%d\n", unsafe.Sizeof("ABCDEFGHIJKLMNOPQRSTUVWXYZ"))

	fmt.Printf("size of chan int: 	%d\n", unsafe.Sizeof(make(chan int)))
	fmt.Printf("size of chan int64:	%d\n", unsafe.Sizeof(make(chan int64)))

	fmt.Printf("size of []int: 		%d\n", unsafe.Sizeof([]int{}))
	fmt.Printf("size of []int{9x}: 	%d\n", unsafe.Sizeof([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	fmt.Printf("size of [10]int{}: 	%d\n", unsafe.Sizeof(make([]int, 10)))

	fmt.Printf("size of map[s]s: 	%d\n", unsafe.Sizeof(map[string]string{}))
	fmt.Println()

	fmt.Printf("size of struct A: 	%d\n", unsafe.Sizeof(A{}))
	fmt.Printf("size of struct B: 	%d\n", unsafe.Sizeof(B{}))
	fmt.Printf("size of struct C: 	%d\n", unsafe.Sizeof(C{}))
	fmt.Printf("size of struct D: 	%d\n", unsafe.Sizeof(D{}))
	fmt.Printf("size of struct E: 	%d\n", unsafe.Sizeof(E{}))
	fmt.Printf("size of struct F: 	%d\n", unsafe.Sizeof(F{}))
	fmt.Printf("size of struct G: 	%d\n", unsafe.Sizeof(G{}))
	fmt.Printf("size of struct H: 	%d\n", unsafe.Sizeof(H{}))
	fmt.Printf("size of struct I: 	%d\n", unsafe.Sizeof(I{}))
	fmt.Printf("size of struct J: 	%d\n", unsafe.Sizeof(J{}))
	fmt.Printf("size of struct K: 	%d\n", unsafe.Sizeof(K{}))
	fmt.Printf("size of struct L: 	%d\n", unsafe.Sizeof(L{}))
	fmt.Printf("size of struct M: 	%d\n", unsafe.Sizeof(M{}))
	fmt.Printf("size of struct N: 	%d\n", unsafe.Sizeof(N{}))
	fmt.Printf("size of struct O: 	%d\n", unsafe.Sizeof(O{}))
	fmt.Printf("size of struct P: 	%d\n", unsafe.Sizeof(P{}))
	fmt.Printf("size of struct Q: 	%d\n", unsafe.Sizeof(Q{}))
	fmt.Printf("size of struct R: 	%d\n", unsafe.Sizeof(R{}))
	fmt.Printf("size of struct S: 	%d\n", unsafe.Sizeof(S{}))
	fmt.Printf("size of struct T: 	%d\n", unsafe.Sizeof(T{}))
	fmt.Printf("size of struct U: 	%d\n", unsafe.Sizeof(U{}))
	fmt.Printf("size of struct V: 	%d\n", unsafe.Sizeof(V{}))

	fmt.Println()
	fmt.Printf("size of Part1: 		%d\n", unsafe.Sizeof(Part1{}))
	fmt.Printf("size of Part2: 		%d\n", unsafe.Sizeof(Part2{}))
}
