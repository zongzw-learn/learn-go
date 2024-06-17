package main

import "log"

type AbstractFactory interface {
	CreateA()
	CreateB()
}

type Factory1 struct{}
type Factory2 struct{}

func (f *Factory1) CreateA() {
	log.Println("A1")
}
func (f *Factory1) CreateB() {
	log.Println("B1")
}
func (f *Factory2) CreateA() {
	log.Println("A2")
}
func (f *Factory2) CreateB() {
	log.Println("B2")
}

var (
	// check type struct implemented all interface's methods.
	_ AbstractFactory = (*Factory1)(nil)
)

/**
 * 简单工厂：就是用同一个工厂类，根据要求（if-else）生产不同的实例。
 * 抽象工厂：抽象工厂定义框架，用具体的工厂类 去生产 具体的实例。
 * 工厂方法：和抽象工厂没有本质的差别，只是，具体工厂仅用于创建单个具体实例。
 */
func main() {
	var f AbstractFactory

	f = &Factory1{}
	f.CreateA()
	f.CreateB()

	f = &Factory2{}
	f.CreateA()
	f.CreateB()

}
