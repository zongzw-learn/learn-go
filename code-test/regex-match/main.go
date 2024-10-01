package main

func main() {

}

func isMatch(s string, p string) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	}

	if 'a'<= p[0] && p[0] <= 'z' {
		if len(p) >1 && p[1] == '*' {
			m1 := isMatch(s, p[2:])
			for i :=0; i<len(s); i++  {
				if isMatch(s, p)
			}
		} else if p[0] == s[0]{
			return isMatch(s[1:], p[1:]) 
		} else {
			return false
		}
	}
}