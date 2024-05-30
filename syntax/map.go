package main

import (
	"fmt"
	"reflect"
)

func main() {
	headers := map[string]string{
		"content-type": "application/json",
		"sessionid":    "3242342352423",
	}

	for k, v := range headers {
		fmt.Printf("%s: %s\n", k, v)
	}

	// 使用reflect 获取map 对象的所有key值。
	keys := reflect.ValueOf(headers).MapKeys()
	fmt.Printf("keys: %v\n", keys)
}
