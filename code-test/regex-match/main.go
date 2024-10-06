package main

import "fmt"

type SP struct {
	s string
	p string
	e bool
}

func main() {
	// s := "aab"
	// p := "c*a*b"

	// s := "mississippi"
	// p := "mis*is*ip*."
	// s := "ab"
	// p := ".*"

	// s := "aaa"
	// p := "ab*a"

	// s := "a"
	// p := "ab*"

	ts := []SP{
		{"aab", "c*a*b", true},
		{"mississippi", "mis*is*ip*.", true},
		{"ab", ".*", true},
		{"aaa", "ab*a", false},
		{"ab", ".*c", false},
		{"aaa", "aaaa", false},
		{"a", "ab*", true},
		{"a", ".*..a*", false},
		{"", ".a*", false},
		{"aaa", "a*a", true},
		{"a", ".*..", false},
		{"bab", "..*", true},
	}
	for i := range ts {
		fmt.Printf("%v(%v) %s %s\n", isMatch(ts[i].s, ts[i].p), ts[i].e, ts[i].s, ts[i].p)
	}
}

func isMatch(s string, p string) bool {
	if p == ".*" {
		return true
	}

	i, j := 0, 0
	for j < len(p) {
		switch p[j] {
		case '.':
			if j+1 < len(p) {
				if p[j+1] == '*' {
					for ; i < len(s); i++ {
						if isMatch(s[i:], p[j+2:]) {
							return true
						}
					}
					return isMatch(s[i:], p[j+2:])
				} else if i < len(s) {
					i++
					j++
				} else {
					return false
				}
			} else {
				if i < len(s) {
					i++
					j++
				} else {
					return false
				}
			}
		case '*':
		default:
			last := p[j]
			if j+1 < len(p) {
				if p[j+1] == '*' {
					for ; i < len(s) && s[i] == last; i++ {
						match := isMatch(s[i:], p[j+2:])
						if match {
							return true
						}
					}
					return isMatch(s[i:], p[j+2:])
				} else {
					if i < len(s) && s[i] == p[j] {
						i++
						j++
					} else {
						return false
					}
				}
			} else {
				if i < len(s) {
					if s[i] == p[j] {
						i++
						j++
					} else {
						return false
					}
				} else {
					return false
				}
			}
		}
		if i >= len(s) {
			return isMatch(s[len(s):], p[j:])
		}
	}
	if i == len(s) && j == len(p) {
		return true
	}
	return false
}
