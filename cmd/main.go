package main

import (
	"net/http"
	"passportScanner/scannersdk"
	"passportScanner/views/pages"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()


    scanner := scannersdk.NewScannerSDK()

	r.GET("/", func(c *gin.Context) {
		pages.Main().Render(c, c.Writer)
	})

	r.POST("/get-document-data", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, "File upload (error): %v", err)
			return
		}

		document, err := scanner.GetDocumentData(file)

		if err != nil {
			c.String(http.StatusBadRequest, "%v", err)
			return
		}

		pages.Response(document).Render(c, c.Writer)
	})

	r.Run(":4000")
}
