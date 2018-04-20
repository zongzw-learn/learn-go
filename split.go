package main

import (
	"fmt"
	"strings"
)

func main() {

	elems := strings.Split("/login|GET,POST", "|")

	fmt.Println(elems)
	elems = strings.Split("/login", "|")
	fmt.Println(elems)

	if len(elems) == 1 {
		elems = append(elems, "a")
	}
	fmt.Println(elems)

}
