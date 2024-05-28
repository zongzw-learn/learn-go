package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// type INT int
// type A struct {
// 	M int
// }

// type B struct {
// 	M INT
// }

type C struct {
	M int
}

type Interface interface {
	DoSomething()
}

func DynamicImplementor(obj interface{}) Interface {
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Struct {
		panic("Object must be a struct")
	}
	// 检查并实现接口
	// ...
	return obj.(Interface)
}

func (c *C) DoSomething() {

}

type Tunnel struct {
	data *byte
}

type Transfer interface {
	Doit()
}
type ToTransfer struct {
	Fd int
	S  string
}

func (tt *ToTransfer) Doit() {

}

var (
	_ Transfer = (*ToTransfer)(nil)
)

func main() {
	// a := A{M: 10}
	// b := (*B)(unsafe.Pointer(&a))
	// fmt.Printf("a: %v, b: %v\n", a, b)

	// // t := reflect.TypeFor[C]()
	// t := reflect.TypeOf(C{})
	// fmt.Printf("t: %v\n", t)
	// // x := t(a)
	// x := DynamicImplementor(&C{M: 10})
	// x.DoSomething()

	t := Tunnel{
		data: (*byte)(unsafe.Pointer(&ToTransfer{Fd: 1, S: "2"})),
	}

	tt := (*ToTransfer)(unsafe.Pointer(t.data))
	fmt.Printf("tt: %d, %s\n", tt.Fd, tt.S)

	ttt := (Transfer)(tt)
	ttt.Doit()

	tttt := (Transfer)(unsafe.Pointer(t.data)) // 本质上这个情况就不可能，因为强行转换后的interface 不具备具化类型，不知道要调用哪个Doit。

}
