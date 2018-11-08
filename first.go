package main

import (
	"net/http"
	yaag_gin "github.com/masato25/yaag/gin"
	"github.com/masato25/yaag/yaag"
	"github.com/gin-gonic/gin"
)
func main() {
	routes := gin.Default()
	yaag.Init(&yaag.Config{
		On:       true,
		DocTitle: "Gin",
		DocPath:  "1111.html",
		BaseUrls: map[string]string{"Production": "", "Staging": "/api/v1"},
	})
	routes.Use(yaag_gin.Document())
	// use other middlewares ...
	routes.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	routes.Run(":8080")
}