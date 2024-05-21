package main

import (
    "fmt"
    "math/rand"
    "time"
)

type Cell struct {
    cache []int
}

var rateRand float64
var total int

func (c *Cell) Next() int {
    fmt.Printf("cell info: %v ", c)
    defer fmt.Printf("\n")

    rand.Seed(int64(time.Now().Nanosecond()))

    poss := rand.Intn(100) 
    if poss < int(100 * rateRand) {
        rlt := rand.Intn(total)
        fmt.Printf("rand -> %v, ", rlt)
        return rlt
    } else {
        if len(c.cache) == 0 {
            rlt := rand.Intn(total)
            fmt.Printf("lean -> rand -> %v, ", rlt)
            return rlt
        } else {
            rlt := rand.Intn(len(c.cache))
            fmt.Printf("lean -> %v, ", c.cache[rlt])
            return c.cache[rlt]
        }
    }
}

func (c *Cell) AddLean(n int) {
    //fmt.Printf("cell add %v\n", n)
    c.cache = append(c.cache, n)
}

func main() {
    
    rand.Seed(int64(time.Now().Nanosecond()))

    rateRand = 0.2
    total = 30
    //c := Cell{[]int{2, 3, 3, 4, 5, 6, 7, 8, 7}}
    //d := Cell{[]int{}}

    //fmt.Printf("c: %v\n", c)
    //fmt.Printf("rand: %v\n", c.Rand(100))
    //fmt.Printf("rand: %v\n", d.Lean())

    cs := make([]Cell, total)

    for i, _ := range cs {
        n := rand.Intn(10)
        for j:=0; j<n; j++ {
            cs[i].AddLean(rand.Intn(total))
        }
        fmt.Printf("cell %v: %v\n", i, cs[i])
    }

    fmt.Printf("cell: %v\n", cs[total-1])

    index := 0
    for k:=0; k<10; k++ {
        index = cs[index].Next()
    }

}
