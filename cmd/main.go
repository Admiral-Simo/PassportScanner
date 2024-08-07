package main

import (
	"passportScanner/handlers"
	"passportScanner/scannersdk"
	"passportScanner/views/pages"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	scanner := scannersdk.NewScannerSDK()

	r.Static("/public", "./public")
	// pages
	r.GET("/", handlers.PresentationPageHandler)
	r.GET("/main", handlers.MainPageHandler)
	r.POST("/get-document-data", handlers.DocumentDataPageHandler(scanner))
	r.GET("/authenticate", func(c *gin.Context) {
		pages.LoginPage(nil).Render(c, c.Writer)
	})
	r.POST("/upload-history", handlers.UploadHistoryPageHandler(scanner))
	r.GET("/contact", handlers.ContactPageHandler)

	r.Run(":4000")
}
