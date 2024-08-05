package main

import (
	"passportScanner/views/pages"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		pages.Main().Render(c, c.Writer)
	})

	r.Run(":4000")
}
