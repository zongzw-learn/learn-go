package main 

import (
    "fmt"
	"encoding/json"
    "io/ioutil"
)

func main() {
    var data interface{}
    bytes, err := ioutil.ReadFile("test.json")
    if err != nil {
        fmt.Printf("failed to read file.\n")
        return 
    }
    err = json.Unmarshal(bytes, &data)
    if err != nil {
        fmt.Printf("failed to unmarshal data.\n")
    }

    fmt.Println(data)
    fmt.Println(data.(map[string]interface{})["message"].(float64))
    fmt.Println(data.(map[string]interface{})["data"].(map[string]string)["username"].(string))

}
