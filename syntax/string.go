package main

import (
	"fmt"
)

type Str string

type StrList struct {
	Item []Str
}

func main() {

	a := StrList{Item: []Str{"a", "b"}}
	fmt.Printf("%v\n", a.Item[0])
	b := string(a.Item[0])
	fmt.Println(b)

	c := "abcdef"
	[]byte(c)[0] = 'A'
	fmt.Println(c) // The original string cannot be changed
}
