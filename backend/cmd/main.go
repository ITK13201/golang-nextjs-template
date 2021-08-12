package cmd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run() {
	engine := gin.Default()
	ua := ""
	// use middleware
	engine.Use(func(c *gin.Context) {
		ua = c.GetHeader("User-Agent")
		c.Next()
	})
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message":    "hello world",
			"User-Agent": ua,
		})
	})
	engine.Run(":3000")
}
