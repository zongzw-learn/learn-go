package main

import (
    "github.com/gin-gonic/gin"
    "fmt"
    "net/http"
    "ginauth/util"
    "github.com/koding/cache"
    "time"
    "strings"

)

func main() {
    StartTick()

    router := gin.Default()

    router.POST("/login", login)
   
    authorized := router.Group("/")
    authorized.Use(AuthRequired()) 
    {
        authorized.GET("/test", test)
    }
    
    fmt.Println("hello go gin!")
    router.Run(":8080")
}



func test(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "test StatusOK",
        })

}

var cacheTTL = cache.NewMemoryWithTTL(20 * 60 * time.Second)

func StartTick() {
    cacheTTL.StartGC(1 * time.Second)

}


func login(context *gin.Context) {
    grant_type := context.PostForm("grant_type")
    fmt.Println("grant type: %s", grant_type)

    if (grant_type == "apikey") {
        apikey := context.PostForm("apikey")
        servicename := context.PostForm("servicename")
        if (apikey == "" || servicename == "") {
            context.JSON(http.StatusBadRequest, 
                fmt.Sprintf("Required: apikey, servicename, but given: '%s', '%s'.", apikey, servicename))
            return
        }
        fmt.Println("apikey: %s, service name: %s", apikey, servicename)

        var endpoint string
        endpoint = "https://iam.stage1.bluemix.net"
        data := make(map[string]string)
        data["grant_type"] = "urn:ibm:params:oauth:grant-type:apikey"
        data["apikey"] = apikey
        
        token, err := util.NewIAMToken(endpoint, data)
        if err != nil {
            fmt.Println(err.Error())
            context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
            return
        }

        fmt.Println(token.AccessToken)
        fmt.Println(token.DumpDetails())

        // insert cache
        sign := strings.Split(token.AccessToken, ".")[2]
        fmt.Println(sign)
        cacheTTL.Set(sign, token.AccessToken)

        context.JSON(http.StatusOK, gin.H{"access_token": token.AccessToken,})

    } else if (grant_type == "authorization_code") {

    } else { 
        context.JSON(http.StatusBadRequest, fmt.Sprintf("The grant_type '%s' is not allowed.", grant_type))
        return
    } 

}

func AuthRequired() gin.HandlerFunc {
    return func (context *gin.Context) {
        fmt.Println("hello in AuthRequired middleware")

        authHeader := context.GetHeader("Authorization")
        if (authHeader == "") {
            context.JSON(http.StatusUnauthorized, "not authenticated.")
            context.Abort()
            return
        }
        sign := strings.Split(authHeader, ".")[2]
        _, err := cacheTTL.Get(sign)
        if (err != nil) {
            context.JSON(http.StatusUnauthorized, "not authenticated.")
            context.Abort()
        } else {
            context.Next()
        }
    }
}

