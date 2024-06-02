package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(ctx *gin.Context) {
		// Call the HTML method of the Context to render a template
		ctx.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"index.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title": "Home Page",
			},
		)
	})

	router.Run(":8080")
}