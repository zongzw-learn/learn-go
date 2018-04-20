package main

import "fmt"

func main() {
    headers := map[string]string {
        "content-type": "application/json",
        "sessionid": "3242342352423",
    }

    for k, v := range headers {
        fmt.Printf("%s: %s\n", k, v)
    }
}
