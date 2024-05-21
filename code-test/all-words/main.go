package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/aaaton/golem/v4"
	"github.com/aaaton/golem/v4/dicts/en"
)

// 找出参数指定的文件中所有的单词（仅由a-zA-Z字母组成）

func main() {

	if len(os.Args) == 1 {
		fmt.Printf("at least one file should be presented.\n")
		os.Exit(1)
	}

	lemmatizer, err := golem.New(en.New())
	if err != nil {
		panic(err)
	}
	words := map[string]bool{}
	is_sep := func(s byte) bool {
		return !(s >= 'a' && s <= 'z') && !(s >= 'A' && s <= 'Z')
	}
	regwd := regexp.MustCompile("[a-zA-Z]+")
	for _, f := range os.Args[1:] {
		b, err := os.ReadFile(f)
		if err != nil {
			fmt.Printf("failed to read file: '%s': %s\n", f, err.Error())
			os.Exit(1)
		}

		l := len(b)
		for i := 0; i < l; {
			if is_sep(b[i]) {
				i += 1
				continue
			}
			j := i
			for ; j < l && !is_sep(b[j]); j++ {
			}
			w := string(b[i:j])
			if regwd.MatchString(w) {
				word := lemmatizer.Lemma(strings.ToLower(w))
				words[word] = true
			}
			i = j
		}
	}
	// fmt.Printf("words: %v\n", words)

	wl := []string{}
	for k := range words {
		wl = append(wl, k)
	}
	fmt.Printf("%v\n", wl)
	fmt.Printf("len: %d\n", len(wl))
}
