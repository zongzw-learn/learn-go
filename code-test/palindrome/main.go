package main

import "fmt"

// 判断数字是否是 palindrome 数字，即，是否对称
func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}

	a := []int{}
	for i := x; i > 0; i /= 10 {
		a = append(a, i%10)
	}
	l := len(a)
	for i := 0; i < l/2; i++ {
		if a[i] != a[l-1-i] {
			return false
		}
	}
	return true
}

func isPalindrome2(num int) bool {
	t := 0
	n := num
	for ; num > 0; num /= 10 {
		t = num%10 + t*10
	}
	return t == n
}

func isPalindrome3(num int) bool {
	t := 0
	for ; num > t; num /= 10 {
		t = num%10 + t*10
		if t == num {
			return true
		}
	}
	return t == num
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	a := fmt.Sprintf("%d", x)

	l := len(a)
	for i := 0; i < l/2; i++ {
		if a[i] != a[l-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Printf("%v\n", isPalindrome(12321))
	fmt.Printf("%v\n", isPalindrome(-12321))
	fmt.Printf("%v\n", isPalindrome(123421))
	fmt.Printf("2 %v\n", isPalindrome2(12321))
	fmt.Printf("2 %v\n", isPalindrome2(1221))
	fmt.Printf("2 %v\n", isPalindrome2(1))

	fmt.Printf("3 %v\n", isPalindrome3(12321))
	fmt.Printf("3 %v\n", isPalindrome3(1221))
	fmt.Printf("3 %v\n", isPalindrome3(1))
	fmt.Printf("3 %v\n", isPalindrome3(111111))
	fmt.Printf("3 %v\n", isPalindrome3(987789))
	fmt.Printf("3 %v\n", isPalindrome3(212121212))
}
