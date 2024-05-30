package main

import (
	"log"
	"reflect"
)

type A struct {
	a int
	b string
	c bool
}

// 实验reflect的相关函数
func main() {
	typeValue()
	callFunc()
}

func others() {
	// 类型赋值
	// newTValue.Elem().FieldByName(newTTag).Set(tValue)

	// 依据 kind 分支
	// reflect.TypeOf(a).Kind()
	// case reflect.Int:
	// case reflect.String:
}

// 动态调用函数
func callFunc() {
	o := &A{}

	in := []reflect.Value{reflect.ValueOf(1), reflect.ValueOf("xyz"), reflect.ValueOf(true)}
	done := reflect.ValueOf(o).MethodByName("Create").Call(in)
	// 类型断言： done[0].Interface().(bool)
	log.Printf("called Created and returned: %v obj: %v", done[0].Interface().(bool), o)
}
func (o *A) Create(a int, b string, c bool) bool {
	o.a = a
	o.b = b
	o.c = c
	return true
}

// 使用reflect实现对象及成员的类型与数值获取
func typeValue() {
	// 声明一个结构体对象
	x := A{a: 1, b: "xyz", c: true}

	// 获取结构体的类型
	log.Printf("type of x: %T", x)
	log.Printf("type of x: %s", reflect.TypeOf(x))
	log.Printf("type of x: %s", reflect.TypeOf(x).String())

	// 获取结构体对象的值
	log.Printf("value of x: %v", x)
	log.Printf("value of x: %v", reflect.ValueOf(x))
	log.Printf("value of x: %v", reflect.ValueOf(x).String())

	// 获取结构体成员的类型和值
	log.Printf("elem num of x: %d", reflect.TypeOf(x).NumField())
	log.Printf("elem num of x: %d", reflect.ValueOf(x).NumField())
	log.Printf("x elem 0: name: %s type: %s, value: %v",
		reflect.TypeOf(x).Field(0).Name,
		reflect.TypeOf(x).Field(0).Type,
		reflect.ValueOf(x).Field(0))

	// 获取结构体成员的类型和值
	log.Printf("x elem 1: name: %s type: %s, value: %v",
		reflect.TypeOf(x).Field(1).Name,
		reflect.ValueOf(x).Field(1).Type(),
		reflect.ValueOf(x).Field(1))
}

// https://segmentfault.com/a/1190000016230264
