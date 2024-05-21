package main

import (
	"fmt"
    "strings"
    "reflect"
    "strconv"
)

func main() {
    arr1 := []int{1, 2, 3, 4}
    arr2 := []int{5, 6, 7}

    arr1 = append(arr1, *arr2)
    fmt.Printf("arr1+arr2: %v\n", arr1)

    m := map[string]string {"a": "1", "b": "2"}
    b := reflect.ValueOf(m).MapKeys()
    fmt.Println(b)
	a := []string{"1", "2", "3", "4", "5", "6"}
	c := []int{1, 2, 3, 4, 5, 6}
    cstr := []string{}
    for _, n := range(c) {
        cstr = append(cstr, strconv.Itoa(n))
    }
    fmt.Println(strings.Join(cstr, ":"))

	fmt.Println(strings.Join(a, ":"))
	s1 := []string{"/login", "/createid/:id", "/create|GET", "/access"}
	s2 := []string{"/login", "/create|GET", "/zongzw"}

	fmt.Printf("%v\n", minus(s1, s2))
	fmt.Printf("%v\n", plus(s1,s2))

}

func minus(a, b []string) []string {
	var result []string
	for _, n := range a {
		found := false
		for _, m := range b {
			if m == n {
				found = true
				break
			}
		}
		if !found {
			result = append(result, n)
		}
	}

	return result
}

func plus(a, b []string) []string {
	var result []string

	for _, n := range a {
		found := false
		for _, m := range result {
			if m == n {
				found = true
				break
			}
		}
		if !found {
			result = append(result, n)
		}
	}

	for _, n := range b {
		found := false
		for _, m := range result {
			if m == n {
				found = true
				break
			}
		}
		if !found {
			result = append(result, n)
		}
	}

	return result
}

