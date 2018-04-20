package main

import (
    "fmt"

    "github.ibm.com/edge/iamutils"
)

func main() {
    iam, err := iamutils.NewClient("activity-tracker", "apikey", "", "", nil)
    if err != nil {
        fmt.Println(err.Error())
    }

    ok, err := iam.Authorize("action", "token")
    if err != nil {
        fmt.Println(err.Error())
    }

    fmt.Println(ok)
}
