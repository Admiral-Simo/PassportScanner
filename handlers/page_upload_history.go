package handlers

import (
	"net/http"
	"passportScanner/scannersdk"
	"passportScanner/views/pages"

	"github.com/gin-gonic/gin"
)

func UploadHistoryPageHandler(scanner scannersdk.ScannerSDK) gin.HandlerFunc {
	return func(c *gin.Context) {
		history, err := scanner.GetUploadHistory()
		// convert from map[string][]string to []struct{date: string, images: []string}

		if err != nil {
			c.String(http.StatusInternalServerError, "Something went wrong", err)
		}

		pages.UploadHistory(history).Render(c, c.Writer)
	}
}
