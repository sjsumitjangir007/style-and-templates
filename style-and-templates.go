package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	// router.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")
	router.LoadHTMLGlob("templates/**/*.tmpl")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index/index.tmpl", gin.H{
			"title": "Welcome",
		})
	})
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "not-found/index.tmpl", gin.H{"title": "Page Not Found", "message": "Page not found"})
	})
	router.Run(":8080")
}
