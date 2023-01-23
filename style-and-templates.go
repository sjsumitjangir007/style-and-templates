package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin/v2"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

func main() {
	router := gin.Default()
	router.Use(apmgin.Middleware(router))
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	router.Static("/assets", "./assets")
	// router.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")
	router.LoadHTMLGlob("templates/**/*.tmpl")
	router.StaticFS("/static", http.Dir("./static"))
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index/index.tmpl", gin.H{
			"title": "SJ, Style & Template Lib.",
		})
	})
	router.GET("/cards", func(c *gin.Context) {
		c.HTML(http.StatusOK, "cards/index.tmpl", gin.H{
			"title": "SJ, Style & Template Lib.",
		})
	})
	router.GET("/badges", func(c *gin.Context) {
		c.HTML(http.StatusOK, "badges/index.tmpl", gin.H{
			"title": "SJ, Style & Template Lib.",
		})
	})
	router.GET("/buttons", func(c *gin.Context) {
		c.HTML(http.StatusOK, "buttons/index.tmpl", gin.H{
			"title": "SJ, Style & Template Lib.",
		})
	})
	router.GET("/forms", func(c *gin.Context) {
		c.HTML(http.StatusOK, "forms/index.tmpl", gin.H{
			"title": "SJ, Style & Template Lib.",
		})
	})
	router.GET("/inputs", func(c *gin.Context) {
		c.HTML(http.StatusOK, "inputs/index.tmpl", gin.H{
			"title": "SJ, Style & Template Lib.",
		})
	})
	router.GET("/lists", func(c *gin.Context) {
		c.HTML(http.StatusOK, "lists/index.tmpl", gin.H{
			"title": "SJ, Style & Template Lib.",
		})
	})
	router.GET("/tabs", func(c *gin.Context) {
		c.HTML(http.StatusOK, "tabs/index.tmpl", gin.H{
			"title": "SJ, Style & Template Lib.",
		})
	})
	router.GET("/datepickers", func(c *gin.Context) {
		c.HTML(http.StatusOK, "datepickers/index.tmpl", gin.H{
			"title": "SJ, Style & Template Lib.",
		})
	})
	router.GET("/dropdowns", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dropdowns/index.tmpl", gin.H{
			"title": "SJ, Style & Template Lib.",
		})
	})
	router.GET("/progressbars", func(c *gin.Context) {
		c.HTML(http.StatusOK, "progressbars/index.tmpl", gin.H{
			"title": "SJ, Style & Template Lib.",
		})
	})
	router.GET("/ratings", func(c *gin.Context) {
		c.HTML(http.StatusOK, "ratings/index.tmpl", gin.H{
			"title": "SJ, Style & Template Lib.",
		})
	})
	router.GET("/skeletons", func(c *gin.Context) {
		c.HTML(http.StatusOK, "skeletons/index.tmpl", gin.H{
			"title": "SJ, Style & Template Lib.",
		})
	})
	router.GET("/spinners", func(c *gin.Context) {
		c.HTML(http.StatusOK, "spinners/index.tmpl", gin.H{
			"title": "SJ, Style & Template Lib.",
		})
	})
	router.GET("/steppers", func(c *gin.Context) {
		c.HTML(http.StatusOK, "steppers/index.tmpl", gin.H{
			"title": "SJ, Style & Template Lib.",
		})
	})
	router.GET("/text", func(c *gin.Context) {
		c.HTML(http.StatusOK, "text/index.tmpl", gin.H{
			"title": "SJ, Style & Template Lib.",
		})
	})
	router.GET("/timelines", func(c *gin.Context) {
		c.HTML(http.StatusOK, "timelines/index.tmpl", gin.H{
			"title": "SJ, Style & Template Lib.",
		})
	})
	router.GET("/tables", func(c *gin.Context) {
		c.HTML(http.StatusOK, "tables/index.tmpl", gin.H{
			"title": "SJ, Style & Template Lib.",
		})
	})
	router.GET("/toasts", func(c *gin.Context) {
		c.HTML(http.StatusOK, "toasts/index.tmpl", gin.H{
			"title": "SJ, Style & Template Lib.",
		})
	})
	router.GET("/videos", func(c *gin.Context) {
		c.HTML(http.StatusOK, "videos/index.tmpl", gin.H{
			"title": "SJ, Style & Template Lib.",
		})
	})
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "not-found/index.tmpl", gin.H{"title": "Page Not Found", "message": "Page not found"})
	})
	router.Run(":8080")
}
