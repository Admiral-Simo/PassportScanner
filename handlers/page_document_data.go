package handlers

import (
	"net/http"
	"passportScanner/scannersdk"
	"passportScanner/views/pages"

	"github.com/gin-gonic/gin"
)

func DocumentDataPageHandler(scanner scannersdk.ScannerSDK) gin.HandlerFunc {
	return func(c *gin.Context) {
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

		pages.DocumentData(document).Render(c, c.Writer)
	}
}
