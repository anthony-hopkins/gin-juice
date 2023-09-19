package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())

	// Configure static resource location
	r.Static("/vendor", "./static/vendor")

	// Tell gin where our templates are stored. 
	r.LoadHTMLGlob("templates/**/**")

	// Set the default root route for the URL
	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "views/index.html", gin.H{
			"title": "Hello Gin!",
		})
	})
	// More reoutes here baseds on strings/regex patterns

	// Start the server providing a console confirmation
	log.Println("Server started!")
	r.Run() // Default port is 8080
}
