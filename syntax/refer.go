package main
import "fmt"

func print(a int) {
    fmt.Printf("value: %d\n", a)
}
func main() {

    res := map[string]string{}
    defer fmt.Println(res)

    res["key"] = "zong"
    return
}

