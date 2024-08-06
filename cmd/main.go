package main

import (
	"net/http"
	"passportScanner/views/pages"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		pages.Main().Render(c, c.Writer)
	})

	r.POST("/get-document-data", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, "File upload (error): %v", err)
			return
		}

		pages.Response(file.Filename).Render(c, c.Writer)
	})

	r.Run(":4000")
}
