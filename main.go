package main

import "github.com/gin-gonic/gin"
import "os"

type User struct {
	Name string
	Age int
	Note string
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/css", "assets/css")

	r.GET("/", func(c *gin.Context) {
		index(c)
	})

	port := os.Getenv("PORT")
  if len(port) == 0 {
      port = "3000"
  }
  r.Run(":" + port)
}

func index(c *gin.Context) {
	users := []User{User{"abc", 21, ""}, User{"def", 34, "xcv"}, User{"ghi", 8, "12489"}}
	c.HTML(200, "index.tmpl", gin.H{
		"title": "Hello, world",
		"users": users,
	})
}
