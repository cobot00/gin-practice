package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	Name string
	Age  int
	Note string
}

type Sample struct {
	ID          uint
	IntValue    int
	RealValue   float32
	StringValue string
}

var db *gorm.DB

func main() {
	db = connectDb()

	r := createRouter()

	setRoute(r)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "5000"
	}
	r.Run(":" + port)
}

func createRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/css", "assets/css")
	return r
}

func setRoute(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		index(c)
	})

	router.GET("/sub", func(c *gin.Context) {
		sub(c)
	})

	router.POST("/post_test", func(c *gin.Context) {
		postTest(c)
		c.Request.URL.Path = "/sub"
		c.Request.Method = "GET"
		router.HandleContext(c)
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
	samples := []Sample{}
	db.Table("sample").Order("id").Find(&samples)
	log.Printf("Sample count: %v", len(samples))

	c.HTML(http.StatusOK, "sample.tmpl", gin.H{
		"title":   "This is Sub",
		"samples": samples,
	})
}

func postTest(c *gin.Context) {
	text1 := c.PostForm("text1")
	number1 := c.PostForm("number1")

	log.Println("text1: " + text1)
	log.Println("number1: " + number1)

	intValue, err := strconv.Atoi(number1)
	if err != nil {
		return
	}

	sample := Sample{StringValue: text1, IntValue: intValue}
	db.Table("sample").Create(&sample)
}

func connectDb() *gorm.DB {
	configs := map[string]string{}
	configs["host"] = os.Getenv("DB_HOST")
	configs["port"] = os.Getenv("DB_PORT")
	configs["user"] = os.Getenv("DB_USER")
	configs["password"] = os.Getenv("DB_PASSWORD")
	configs["dbname"] = os.Getenv("DB_SCHEMA")
	configs["sslmode"] = os.Getenv("DB_SSL")

	buf := []string{}
	for k, v := range configs {
		buf = append(buf, k+"="+v)
	}
	params := strings.Join(buf, " ")

	db, err := gorm.Open("postgres", params)
	if err != nil {
		log.Println("DB connect error!!")
		log.Println(err)
		return nil
	}

	log.Println("DB connect success!")

	return db
}
