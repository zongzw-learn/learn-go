package main
import (

	"fmt"
	
)

func main() {

	a := []int64{1, 2, 3, 4, 5}
	b := []int64{2, 3, 4, 5}

	for _, val := range a {
		fmt.Println(val)
	}

	for _, val := range b {
		fmt.Println(val)
	}

}

