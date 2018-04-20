package main

import (
    "fmt"
    "reflect"
)
func main() {
    arg1 := 1
    arg2 := 2
    fmt.Printf("%d\n", acall(func1, arg1, arg2))
    return 
}

func acall(afunc func(...int) int, arg ...int) int {
    a := afunc(arg...)
    return a
}
func func1(arg ... int) int {
    fmt.Printf("type of arg:%v\n", reflect.TypeOf(arg))
    return  arg[0] + arg[1]
}

