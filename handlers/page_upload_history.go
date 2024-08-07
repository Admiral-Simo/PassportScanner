package handlers

import (
	"errors"
	"net/http"
	"passportScanner/scannersdk"
	"passportScanner/views/pages"

	"github.com/gin-gonic/gin"
)

const user = "root"
const pass = "gmsoft3344"

func UploadHistoryPageHandler(scanner scannersdk.ScannerSDK) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		if username != user || password != pass {
			pages.LoginPage(errors.New("invalid email or password.")).Render(c, c.Writer)
			return
		}
		history, err := scanner.GetUploadHistory()
		// convert from map[string][]string to []struct{date: string, images: []string}

		if err != nil {
			c.Status(http.StatusInternalServerError)
			pages.ErrorPage("Something went wrong.").Render(c, c.Writer)
			return
		}

		pages.UploadHistory(history).Render(c, c.Writer)
	}
}
