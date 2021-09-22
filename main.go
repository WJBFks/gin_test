package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/*")
	r.StaticFS("/static", http.Dir("./static"))
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!222222")
	})
	r.GET("/index.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.GET("/intro.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "intro.html", gin.H{})
	})
	r.GET("/start.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "start.html", gin.H{})
	})
	r.Run(":8000")
}
