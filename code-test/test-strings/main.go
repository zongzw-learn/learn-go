package main

import "fmt"

// 找出字符串上最大的不重复子串
func main() {
	// s := "abababcabcabcdabcdeabcabcdefabcab"
	s := "We can access the element in a golang map using its key. if the key is not present, it will return the value type zero value."
	fmt.Printf("%s\n", s[0:])

	mi := 0
	ml := 0
	h := map[byte]int{}
	for i := 0; i < len(s); i++ {
		if l, found := h[s[i]]; found {
			for k, v := range h {
				if v <= l {
					delete(h, k)
				}
			}
			h[s[i]] = i
		} else {
			h[s[i]] = i
			if ml < len(h) {
				ml = len(h)
				mi = i
			}
		}
	}

	fmt.Printf("%s\n", s[mi-ml+1:mi+1])
}
