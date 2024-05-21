package main

import (
    "fmt"
)
type Abser interface {
    Abs() float64
}

type Myfloat64 float64
func (v Myfloat64) Abs() float64 {
    return 2*float64(v)
}

func (v *Vertex) Abs() float64 {
    return v.x*v.x + v.y*v.y
}

type Vertex struct {
    x float64
    y float64
}
func main() {

    var a Abser
    v1 := Myfloat64(45.3)
    v2 := Vertex{23.4, 56.1}

    a = v1
    fmt.Printf("%v\n", a.Abs())

    a = &v2
    fmt.Printf("%v\n", a.Abs())
}
