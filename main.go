package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

type User struct {
	Name string
	Age  int
	Note string
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/css", "assets/css")

	setRoute(r)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "5000"
	}
	r.Run(":" + port)
}

func setRoute(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		index(c)
	})

	router.GET("/sub", func(c *gin.Context) {
		sub(c)
	})
}

func index(c *gin.Context) {
	users := []User{{"abc", 21, ""}, {"def", 34, "xcv"}, {"ghi", 8, "12489"}}
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Hello, world",
		"users": users,
	})
}

func sub(c *gin.Context) {
	users := []User{{"hoge", 21, ""}, {"fuga", 34, "xcv"}, {"piyo", 8, "12489"}}
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "This is Sub",
		"users": users,
	})
}
