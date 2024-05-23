package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// fmtScanf()
	// fmtScanln()
	// bufioReadString()
	// bufioRead()
	fmtSscanx()
}

func fmtSscanx() {
	var name string
	var alphabet_count int

	// Calling the Sscan() function which
	// returns the number of elements
	// successfully scanned and error if
	// it persists
	n, err := fmt.Sscan("GFG 3", &name, &alphabet_count)

	// Below statements get executed if there is any error
	if err != nil {
		panic(err)
	}

	// Printing the number of elements and each elements also
	fmt.Printf("%d: %s, %d\n", n, name, alphabet_count)

	// Calling the Sscanf() function which
	// returns the number of elements
	// successfully parsed and error if
	// it persists
	n, err = fmt.Sscanf("GFG is having 3 alphabets.",
		"%s is having %d alphabets.", &name, &alphabet_count)

	// Below statements get
	// executed if there is any error
	if err != nil {
		panic(err)
	}

	// Printing the number of
	// elements and each elements also
	fmt.Printf("%d: %s, %d\n", n, name, alphabet_count)
}

func bufioRead() {
	b := make([]byte, 200)
	fmt.Printf("input string of length < 200:")
	reader := bufio.NewReader(os.Stdin)
	n, err := reader.Read(b)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	// 注意，b中存储的是 "abcdf\n\x00\x00\x00..."，所以trim并不起作用
	// fmt.Printf("read %d: %s\n", n, strings.Trim(string(b), "\n"))
	fmt.Printf("read %d: %s\n", n, b[0:n-1])
}

func bufioReadString() {
	fmt.Printf("read from stdin with bufio.ReadString: ")
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	fmt.Printf("read: %s\n", strings.Trim(str, "\n"))
}

// 使用fmt.Scanln 以回车为分隔符输入内容到指定变量
func fmtScanln() {
	var a int
	var s string
	var b bool
	fmt.Printf("read from stdin with fmt.Scanln: \n")

	fmt.Printf("input num:")
	fmt.Scanln(&a)
	fmt.Printf("input string:")
	fmt.Scanln(&s)
	fmt.Printf("input bool:")
	fmt.Scanln(&b)

	fmt.Printf("%d %s %t\n", a, s, b)
}

// 使用fmt.Scanf 按照空格逐个从stdin读入单词写入到指定变量，分隔符必须为空格
func fmtScanf() {
	var a int
	var s string
	var b bool
	fmt.Printf("read from stdin with fmt.Scanf: %%d %%s %%t\n")
	n, err := fmt.Scanf("%d %s %t", &a, &s, &b)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("done reading from stdin %d, %d %s %t\n", n, a, s, b)
}
