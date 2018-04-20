package main

import (
	"fmt"
	"time"
	"errors"
	"regexp"
	"net/http"
	"os"
	"strings"
	"encoding/json"

	"github.com/go-xorm/xorm"
	"github.com/gin-gonic/gin"
	"github.com/koding/cache"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {

	Username string
	Password string

}

type Article struct {

	Title string
	Pages int64
	Author string
}

type Authorization struct {
	ApiKey string 	`json:"apikey" binding:"required"`
	Role string 	`json:"role" binding:"required"`
	Service string 	`json:"service" binding:"required"`
	Tags []string	`json:"tags" binding:"required"`
}

type RoleOps struct {
	Role string 			`json:"role" binding:"required"`
	AllowedAPIs []string	`json:"allowedapis" binding:"required"`
}


type AccessConfig struct {
	AuthInfo []Authorization `json:"authorization"`
	ROInfo []RoleOps 		 `json:"roleops"`
}

var driver = "sqlite3"
var dbname = "./test.db"
var CacheTTL = cache.NewMemoryWithTTL(3600 * time.Second)


var orm *xorm.Engine

func InitAuthDB() {
	fmt.Println("hello postgres")
	var err error
	orm, err = xorm.NewEngine(driver, dbname)
	err = orm.Sync(/*new(User), new(Article), */ new(Authorization), new(RoleOps))
	if err != nil {
		fmt.Println(err)
	}
}
func main() {


	EmptyAuthData()
	InitAuthDB()
	InitializeAuthData()

	router := gin.Default()

	router.GET("/access", Access)
	router.Any("/access/:type", Access)

	router.Use(Authorize)

	router.GET("/status", Status)
	router.GET("/create/:id", Create)

	router.POST("/create/:id", Create)
	router.GET("/login", Login)

	router.Run(":8080")
}

func EmptyAuthData() {

	orm, err := xorm.NewEngine(driver, dbname)
	if err != nil {
		fmt.Println(err)
		return 
	}
	if err = orm.DropTables(new(Authorization), new(RoleOps)); err != nil {
		fmt.Printf("Failed to fresh tables: %s\n", err.Error())
		return 
	}

}
func Login(context *gin.Context) {

	context.JSON(http.StatusOK, gin.H{"message": "Request Complete!\n"})

}
func Status(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Request Complete!\n"})
}
func Create(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Request Complete!" + fmt.Sprintf("%d \n", context.Param("id"))})
}

func Authorize(context *gin.Context) {

	//orm, err := xorm.NewEngine(driver, dbname)

	fmt.Println("Authorizing ...")
	//fmt.Printf("hello: %+v\n", context)
	//fmt.Printf("Request: %+v\n", context.Request)
	
	//hello: &{writermem:{ResponseWriter:0xc4202580e0 size:-1 status:200} 
	//Request:0xc4201dc900 Writer:0xc42021e9a0 Params:[] handlers:[0x9b79b0 0x9b86f0 0xa1dcb0 0xa1dbb0] index:2 engine:0xc4202167e0 
	//Keys:map[] Errors: Accepted:[]}
	//Request: &{Method:GET URL:/status Proto:HTTP/1.1 ProtoMajor:1 
	//ProtoMinor:1 Header:map[User-Agent:[curl/7.52.1] Accept:[*/*]] 
	//Body:{} GetBody:<nil> ContentLength:0 TransferEncoding:[] Close:false 
	//Host:localhost:8080 Form:map[] PostForm:map[] MultipartForm:<nil> Trailer:map[] 
	//RemoteAddr:127.0.0.1:36786 RequestURI:/status TLS:<nil> Cancel:<nil> Response:<nil> ctx:0xc4201d2cc0}

	// access_token -> apikey
	access_token := "access_token"
	CacheTTL.Set(access_token, map[string]string{"apikey": "q0IMVSOHDoKYEauQwjyScG4PJQL2_jySgeS8ZkjsFcpx", "service": "Cloudant"})
	cacheData, err := CacheTTL.Get(access_token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Cannot find corresponding userinfo."})
		context.Abort()
		return
	}
	userinfo := cacheData.(map[string]string)
	// service 
	service := userinfo["service"]

	// endpoint
	endpoint := context.Request.URL.Path

	// method
	method := context.Request.Method
	fmt.Printf("method: %s\n", method)

	err = AuthorizeImp(userinfo["apikey"], service, endpoint, method)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		context.Abort()
		return
	}

	fmt.Printf("user authorized: apikey: %s, service: %s, endpoint: %s, method: %s\n", userinfo["apikey"], service, endpoint, method)

}

func InitializeAuthData() {

	orm, err := xorm.NewEngine(driver, dbname)
	if err != nil {
		fmt.Println(err)
		return 
	}

	// read from initial access configuration

	var accessConfig AccessConfig
	configFile, err := os.Open("access.json")
	if err != nil {
		fmt.Printf("failed to open file %s\n", err.Error())
		return 
	}
	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&accessConfig); err != nil {

		fmt.Printf("failed to decode content of file: %s\n", err.Error())
		return 
	}

	//fmt.Printf("data read: %v\n", accessConfig)

	var affected int64
	for _, val := range accessConfig.AuthInfo {
		//fmt.Println(val)
		exists, err := orm.Exist(&val)
		if err != nil {
			fmt.Sprintf("failed to check existing: %v\n", err)
			continue
		}
		//fmt.Printf("exists: %t\n", exists)
		if !exists {
			affected, err = orm.InsertOne(&val)
			if err != nil {

				fmt.Printf("failed to insert test data.:%v\n", val)
			}
			fmt.Printf("insert test data: %v, affected:%d\n", val, affected)
		}
		
	}

	for _, val := range accessConfig.ROInfo {

		//fmt.Println(val)
		exists, err := orm.Exist(&val)
		if err != nil {
			fmt.Sprintf("failed to check existing: %v\n", err)
			continue
		}
		//fmt.Printf("exists: %t\n", exists)
		if !exists {
			affected, err = orm.InsertOne(&val)
			if err != nil {

				fmt.Printf("failed to insert test data.:%v\n", val)
			}	

			//fmt.Printf("insert test data: %v, affected:%d\n", val, affected)
		}
		
	}
	
}

func AuthorizeImp(apikey, service, endpoint, method string) error {
	fmt.Printf("apikey: %s, service: %s, endpoint: %s, method: %s\n", apikey, service, endpoint, method)


	orm, err := xorm.NewEngine(driver, dbname)

	var auth = Authorization{ApiKey: apikey, Service: service}
	has, err := orm.Get(&auth)
	if err != nil || !has {
		return errors.New("No record from Authorization found. ")
	}

	//fmt.Printf("found here: auth: %+v\n", auth)
	var roleops = RoleOps{Role: auth.Role}
	has, err = orm.Get(&roleops)
	if err != nil || !has {
		return errors.New("No record from RoleOps found. ")
	}

	//fmt.Printf("found here: roleops: %+v\n", roleops)

	reParam := regexp.MustCompile(":\\w+")
	var reStr string
	matched := false
	for _, val := range roleops.AllowedAPIs {

		//fmt.Printf("%s\n", val)
		elems := strings.Split(val, "|")
		//fmt.Printf("elems: %v\n", elems)

		reStr = reParam.ReplaceAllString(elems[0], "\\w+")
		matchedEndpoint, err := regexp.MatchString(reStr, endpoint)
		if err != nil {

			fmt.Printf("Failed to match string %s with pattern %s\n", endpoint, reStr)
			return errors.New(fmt.Sprintf("Failed to match string %s with pattern %s\n", endpoint, reStr))
		}

		var matchedMethod bool
		if len(elems) == 1 || strings.Contains(elems[1], method) {
			matchedMethod = true
		}

		if matchedEndpoint && matchedMethod {
			fmt.Printf("endpoint: %s matches %s\n", endpoint, val)
			matched = true
			break
		}
	}

	if !matched {
		return errors.New("Endpoint " + endpoint + " is not included in allowedAPIs")
	}
	
	return nil
}

// Access lets user manage accesses. 
func Access(context *gin.Context) {
	uri := context.Request.URL.Path
	method := context.Request.Method

	if orm == nil {
		context.JSON(http.StatusInternalServerError, 
			gin.H{"message": fmt.Sprintf("database engine turns to nil")})
		return
	}

	if uri == "/access" {
		if method != "GET" {
			context.JSON(http.StatusBadRequest, 
				gin.H{"message": fmt.Sprintf("method '%s' for uri '%s' is not supported.", method, uri)})
			return
		}
		// dump to json and response the result.
		
		var rltAuth []Authorization
		var rltROps []RoleOps

		err := orm.Find(&rltAuth)
		if err != nil {
			context.JSON(http.StatusBadRequest, 
				gin.H{"message": "Failed to get all record"})
			return
		}

		err = orm.Find(&rltROps)
		if err != nil {
			context.JSON(http.StatusBadRequest, 
				gin.H{"message": "Failed to get all record"})
			return
		}		
		
		context.JSON(http.StatusOK, gin.H{"authorization": rltAuth, "roleops": rltROps})
		return 
	} else {
		opType := context.Param("type")
		if opType == "authorization" {
			if method == "GET" {
				context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Use /access endpoint instead to get all. ")})
				return
			} else {
				var val Authorization
				if err := context.ShouldBindJSON(&val); err != nil {
					context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				existVal := Authorization{ApiKey:val.ApiKey, Service: val.Service}
				exists, err := orm.Exist(&existVal)
				if err != nil {
					context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return 
				}

				if method == "POST" {
					if exists {
						affected, err := orm.Table("authorization").Where("api_key = ?", val.ApiKey).And("service = ?", val.Service).Update(val)
						if err != nil {
							context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
							return
						}
						context.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Updated: %v, affected: %d", val, affected)})
						return
					} else {
						affected, err := orm.Insert(val)
						if err != nil {
							context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
							return
						}
						context.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("created: %v, affected: %v", val, affected)})
						return
					}

				} else if method == "DELETE" {
					if exists {
						affected, err := orm.Delete(&Authorization{ApiKey: val.ApiKey, Service: val.Service})
						if err != nil {
							context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
							return
						}
						context.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("deleted: %v, affected: %d", val, affected)})
						return
					} else {
						context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Record %v not exists", val)})
						return
					}
				} else {
					context.JSON(http.StatusBadRequest, 
						gin.H{"message": fmt.Sprintf("method '%s' for uri '%s' is not supported.", method, uri)})
				}
			}
			

		} else if opType == "roleops" {

			if method == "GET" {
				context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Use /access endpoint instead to get all. ")})
				return
			} else {
				var val RoleOps
				if err := context.ShouldBindJSON(&val); err != nil {
					context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				existVal := RoleOps{Role:val.Role}
				exists, err := orm.Get(&existVal)
				if err != nil {
					context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return 
				}

				if method == "POST" {
					if exists {
						changed := plus(val.AllowedAPIs, existVal.AllowedAPIs)
						var valRlt []RoleOps
						affected, err := orm.Where("role = ?", val.Role).Update(&RoleOps{Role: val.Role, AllowedAPIs: changed})
						if err != nil {
							context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
							return
						}
						context.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Updated: %v, affected: %d", val, affected)})
						return
					} else {
						affected, err := orm.Insert(val)
						if err != nil {
							context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
							return
						}
						context.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("created: %v, affected: %v", val, affected)})
						return
					}
				} else if method == "DELETE" {
					if exists {
						changed := minus(existVal.AllowedAPIs, val.AllowedAPIs)
						affected, err := orm.Update(&RoleOps{Role: val.Role, AllowedAPIs: changed})
						if err != nil {
							context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
							return
						}
						context.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Deleted: %v, affected: %d", val, affected)})
						return
					} else {
						context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Record %v not exists", val)})
						return
					}
				} else {
					context.JSON(http.StatusBadRequest, 
						gin.H{"message": fmt.Sprintf("method '%s' for uri '%s' is not supported.", method, uri)})
				}
			}
			
		} else {
			context.JSON(http.StatusBadRequest, 
				gin.H{"message": fmt.Sprintf("type %s not supported.", opType)})
			return 
		}
	}
}

func minus(a, b []string) []string {
	var result []string
	for _, n := range a {
		found := false
		for _, m := range b {
			if m == n {
				found = true
				break
			}
		}
		if !found {
			result = append(result, n)
		}
	}

	return result
}

func plus(a, b []string) []string {
	var result []string

	for _, n := range a {
		found := false
		for _, m := range result {
			if m == n {
				found = true
				break
			}
		}
		if !found {
			result = append(result, n)
		}
	}

	for _, n := range b {
		found := false
		for _, m := range result {
			if m == n {
				found = true
				break
			}
		}
		if !found {
			result = append(result, n)
		}
	}

	return result
}
