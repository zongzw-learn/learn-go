package main

import (
    "fmt"
)

func fibonacci(c, quit chan int) {
    x, y := 0, 1
    for {
        select {
            case c <- x:
                x, y = y, x + y
            case <- quit:
                fmt.Println("quit")
                return 
        }
    }
}

func main() {
    c := make(chan int)
    quit := make(chan int)
    go func() {
    
        for i:= 0; i < 10; i++ {
            fmt.Println(<-c)
        }
        quit <- 0
    }()

    fibonacci(c, quit)

    ii := make(chan int)
    jj := make(chan int)
    ii <- 1
    jj <- 2 
    //jj  <- <- ii

    fmt.Printf("%v, %v \n", <- ii, <- jj)
}


