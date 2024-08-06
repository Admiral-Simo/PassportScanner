package main

import (
	"passportScanner/handlers"
	"passportScanner/scannersdk"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	scanner := scannersdk.NewScannerSDK()

    // pages
	r.GET("/", handlers.MainPageHandler)
	r.POST("/get-document-data", handlers.DocumentDataPageHandler(scanner))
	r.GET("/upload-history", handlers.UploadHistoryPageHandler(scanner))

	r.Run(":4000")
}
