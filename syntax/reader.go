package main

import (
    "fmt"
    "io"
    "strings"
)

func main() {
    r := strings.NewReader("ERR: 2017-07-10 19:19:24.497164555 +0800 CST, it didn't work")

    b := make([]byte, 8)

    for {
        n, err := r.Read(b)
        fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
        fmt.Printf("b[:n] = %q\n", b[:n])
        if err == io.EOF {
            break
        }
    }
    
}
