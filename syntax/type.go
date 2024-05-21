package main

import "fmt"

func main() {

    var i interface{}

    i = 32

    switch value := i.(type) {
        case int: 
            fmt.Printf("int: %d\n", value)
        default: 
            fmt.Printf("%v\n", "not known")
    }

    var j interface{}

    k := map[string]string{"a": "a", "b": "b"}

    j = k

    x := j.(map[string]string)
    fmt.Printf("%s\n", x["a"])
}
