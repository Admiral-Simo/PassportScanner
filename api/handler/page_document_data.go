package handler

import (
	"PassportScanner/scannersdk"
	"PassportScanner/views/pages"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DocumentDataPageHandler(scanner scannersdk.ScannerSDK) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, "File upload (error): %v", err)
			return
		}

		tmpFilePath := "./public/" + file.Filename
		if err := c.SaveUploadedFile(file, tmpFilePath); err != nil {
			c.String(http.StatusInternalServerError, "Error saving file: %v\n", err)
			return
		}

		document, err := scanner.GetDocumentData(file)
		if err != nil {
			c.String(http.StatusBadRequest, "%v", err)
			return
		}

		imageUrl := "/public/" + file.Filename

		pages.DocumentData(document, imageUrl).Render(c, c.Writer)
	}
}
