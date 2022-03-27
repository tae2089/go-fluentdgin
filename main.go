package main

import (
	"github.com/fluent/fluent-logger-golang/fluent"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var logError error
var logger *fluent.Fluent

func init() {
	//logger, logError = fluent.New(fluent.Config{FluentNetwork: "tcp", FluentPort: 24223})
	logger, logError = fluent.New(fluent.Config{FluentNetwork: "tcp"})
	if logError != nil {
		panic(logError.Error())
	}

}

func getting(c *gin.Context) {
	tag := "myapp.log"
	var data = map[string]string{
		"foo":  "bar",
		"hoge": "hoge",
	}
	logger.Post(tag, data)
	c.String(http.StatusOK, "hello world")
}

func postting(c *gin.Context) {
	data := c.PostForm("data")
	log.Println(data)
}

func main() {
	//gin 선언하기
	router := gin.Default()
	//Get Mapping
	router.GET("/", getting)
	router.POST("/", postting)
	//서버 실행
	router.Run(":4001")
}
