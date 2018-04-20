package main

import (
    "fmt"
    "regexp"
)

func main() {
    fmt.Println("hello regex")

    source := "/mb/:magicbox/:customername"

    re := regexp.MustCompile(":\\w+")
    reStr := re.ReplaceAllString(source, "\\w+")
    dest := "/mb/23/3_4"


    matched, err := regexp.MatchString(reStr, dest)
    if err != nil {
    	fmt.Println(err)
    }

    fmt.Printf("%s matched %s: %t\n", reStr, dest, matched)
    
    










}
