package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/test", Test)

	router.Run(":8000")
}

// summary: `form:"hosts"` are needed to bind query
type Spec struct {
	Hosts    []string `json:"hosts,omitempty" form:"hosts"`
	Clusters []string `json:"clusters,omitempty" form:"clusters"`
}

/*
func Query2Struct(query string, spec *interface{}) error {
	kvs := strings.Split(query, "&")
	var kva map[string][]string
	for _, n := range kvs {
		t := strings.Split(n, "=")
		if len(t) != 2 {
			return fmt.Errorf("incorrect format: %s", n)
		}
		k := t[0]
		if val, ok := kva[k]; !ok {
			kva[k] = []string{}
		}
		kva[k] = append(kva[k], t[1])
	}
	fmt.Printf("kva: %v\n", kva)

	err := json.Unmarshal(byte())
}
*/
func Test(c *gin.Context) {
	fmt.Printf("url: %s\n", c.Request.URL)

	var data Spec
	err := c.ShouldBindQuery(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err})
		return
	}
	fmt.Printf("spec: %v\n", data)
	c.JSON(http.StatusOK, gin.H{"message": data})

	return
}
