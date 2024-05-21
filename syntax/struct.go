package main 

import "fmt"

type A struct {
    X int
}

type B struct {
    A
    Y string
}

func main() {
    fmt.Println(B{A:A{1}, Y: "str"})
}
