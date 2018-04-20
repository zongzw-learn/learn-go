package main

import (
    "fmt"
)
type Vertex struct {
    x, y float64
}

func (v Vertex) String() string {

    return fmt.Sprintf("(%v, %v)", v.x, v.y)
}

func main() {

    v := Vertex{23.4, 76.4}

    fmt.Printf("%v\n", v)
}
