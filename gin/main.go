package main

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/sujeetdpa/go_learning/gin/controller"

	"github.com/gin-gonic/gin"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	// router := gin.Default() //Default initialisation, Default logger is used which is console

	//Our own configuratino
	router := gin.New()
	router.Use(gin.Logger())

	//Adding basic auth to admin
	adminAuth := gin.BasicAuth(gin.Accounts{
		"admin": "admin",
	})

	router.GET("/hello", handleHello)
	router.POST("/post", handlePost)
	router.GET("/query-param", handleQueryParam)
	router.GET("/path-variable/name/:name/age/:age", handlePathVariable)

	//Router Grouping
	adminRouter := router.Group("/admin", adminAuth)
	{
		adminRouter.GET("/greet", controller.HandleAdminGreet)
	}

	clientRouter := router.Group("/client")
	{
		clientRouter.GET("/greet", controller.HandleClientGreet)
	}

	// router.Run() // This will by default create a server and listen on port 8080
	// http.ListenAndServe(":8080", router) //This is same as router.Run()

	//Lets create a custom server
	//We can pass custom server properties
	server := http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server.ListenAndServe()
}
func handleHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello Go lang modular",
	})
}
func handlePost(c *gin.Context) {
	//Reading body
	body := c.Request.Body
	value, err := io.ReadAll(body)
	var student Student
	json.Unmarshal(value, &student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":   "Unaable to read body",
			"timestamp": time.Now(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "post handeled properly",
		"body":    student,
	})
}
func handleQueryParam(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello Go lang modular",
		"name":    name,
		"age":     age,
	})
}
func handlePathVariable(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello Go lang modular",
		"name":    name,
		"age":     age,
	})
}
