package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())

	r.Static("/vendor", "./static/vendor")

	r.LoadHTMLGlob("templates/**/**")

	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "views/index.html", gin.H{
			"title": "Hello Gin!",
		})
	})
	log.Println("Server started!")
	r.Run() // Default port is 8080
}
