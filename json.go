package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Authorization struct {
	ApiKey  string   `json:"apikey" binding:"required"`
	Role    string   `json:"role" binding:"required"`
	Service string   `json:"service" binding:"required"`
	Tags    []string `json:"tags" binding:"required"`
}

type RoleOps struct {
	Role        string   `json:"role" binding:"required"`
	AllowedAPIs []string `json:"allowedapis" binding:"required"`
}

type AccessConfig struct {
	AuthInfo []Authorization `json:"authorization"`
	ROInfo   []RoleOps       `json:"roleops"`
}

type CloneVMStruct struct {
	Src         string `json:"src" binding:"required"`
	Dst         string `json:"dst" binding:"required"`
	Datacenter  string `json:"datacenter" binding:"required"`
	HostCluster string `json:"hostcluster" binding:"required"`
	Folder      string `json:"folder,omitempty" binding:"optional"`
	Pool        string `json:"pool" binding:"required"`
}

func main() {

	cloneparam := `
	{
		"sc": "srcvm",
		"dst": "dstvm",
		"datacenter": "Landing",
		"hostcluster": "Basic",
		"pool": "Resource"
	}
	`

	var s CloneVMStruct

	clerr := json.Unmarshal([]byte(cloneparam), &s)
	if clerr != nil {
		fmt.Printf("failed to unmarshal string %s\n", clerr.Error())
		return
	}

	fmt.Printf("struct: %v\n", s)
    data, _ := json.Marshal(s)
    fmt.Printf("json: %v\n", string(data))
	return

	var accessConfig AccessConfig
	configFile, err := os.Open("/go/src/pg-test/access.json")
	if err != nil {
		fmt.Printf("failed to open file %s\n", err.Error())
		return
	}
	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&accessConfig); err != nil {

		fmt.Printf("failed to decode content of file: %s\n", err.Error())
		return
	}

	fmt.Printf("data read: %v\n", accessConfig)

	var dat map[string]interface{}

	jsonStr := "{\"userid\": \"bjzongzw@cn.ibm.com\", \"role\": \"admin\", \"service\": \"bluemix\", \"tags\": [\"newcreated\"], \"userinfo\": {\"id\": 2342342324, \"company\": \"elsewhere of ibm\"}}"
	if err := json.Unmarshal([]byte(jsonStr), &dat); err != nil {

		fmt.Printf("failed to unmarshal the string to json: %s\n", err.Error())
		return
	} else {
		fmt.Printf("dat: %v\n", dat)
		fmt.Printf("data: %v\n", dat["userinfo"])

	}

}
