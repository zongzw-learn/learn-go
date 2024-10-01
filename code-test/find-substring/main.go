package main

import "fmt"

func main() {
	// s := "barfoofoobarthefoobarman"
	s := "foobarthefoobarman"
	words := []string{"foo", "bar", "the"}
	fmt.Println(findSubstring(s, words))
}

func findSubstring(s string, words []string) []int {
	rlt := []int{}
	for i := 0; i < len(s); i++ {
		if doesMatch(s[i:], words) {
			rlt = append(rlt, i)
		}
	}
	return rlt
}

func doesMatch(s string, words []string) bool {
	if len(words) == 0 {
		return true
	}

	if len(words[0]) > len(s) || words[0] != s[0:len(words[0])] {
		return false
	}

	return doesMatch(s[len(words[0]):], words[1:])
}
