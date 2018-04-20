package main

import (
	"fmt"
	"time"
	"math/rand"
)
func main() {
	fmt.Println(time.Now().Format("20060102150405"))
	fmt.Println(rand.Intn(200000))
}
