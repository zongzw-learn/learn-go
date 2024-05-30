package main

import (
	"fmt"
)

type I interface {
	do()
}

type X struct {
	xx I
}

type AA struct {
}

func (aa *AA) do() {}

type BB struct {
}

func (bb *BB) do() {}

// 实验interface{} 类型转换。
func main() {

	// interface{} 类型可以“具化”，通过类型断言的方式将本来的类型还原。
	aa := AA{}

	var allType interface{}
	allType = aa

	// bb, _ := aa.(BB) // invalid operation: aa (variable of type AA) is not an interface
	bb, _ := allType.(BB)
	bb.do()

	ap := &AA{}
	var p interface{}
	p = ap
	cc, _ := p.(I)
	fmt.Println(cc)

	// t := reflect.TypeOf(AA{})
	// tt := AA{}.(type) // use of .(type) outside type switch
	// var a interface{}
	// b, _ := a.(tt) // t (variable of type reflect.Type) is not a type
}
