package main

import (
    "fmt"
    "math"
    "time"
)

type MyError struct {
    when time.Time
    what string
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
    //return fmt.Sprintf("negative number: %v", e)
}

func (e MyError) Error() string {
    return fmt.Sprintf("%v: %v", e.when, e.what)
}

func Sqrt(x float64) (float64, error) {
    if x < 0 {
        //return 0, MyError{time.Now(), "negative number!"}
        return 0, ErrNegativeSqrt(x)
    }
    return math.Sqrt(x), nil
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}

