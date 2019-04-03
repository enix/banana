package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r.Use(cors.Default())

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/buckets", func(c *gin.Context) {
		var o ObjectStorage
		o.Connect("https://object-storage.r1.nxs.enix.io/", credentials.NewStaticCredentials("1ff11a90064045c39ef40f73bd1838fb", "e5a508a5c0b748be8fc78a77a83b8a68", ""))
		list, _ := o.ListBuckets()
		out, _ := json.Marshal(list)
		c.String(http.StatusOK, string(out))
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		fmt.Println("oko")

		user := c.MustGet(gin.AuthUserKey).(string)

		fmt.Println(user)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			fmt.Println("oko2")
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":80")
}
