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
			c.Status(http.StatusInternalServerError)
			pages.ErrorPage().Render(c, c.Writer)
			return
		}

		pages.UploadHistory(history).Render(c, c.Writer)
	}
}
