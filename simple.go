package main 

import (
    "fmt"
)

var i = 23
var fetched = map[string]int{}
func main() {

    fmt.Printf("%v\n", i)
    fmt.Printf("%v\n", 2 << 11 - 1)
    fetched["a"] = 1
    fmt.Printf("%v, %v\n", fetched["a"], fetched["b"])
    
}
