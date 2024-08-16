package handler

import (
	"errors"
	"net/http"
	"PassportScanner/models"
	"PassportScanner/scannersdk"
	"PassportScanner/views/pages"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
)

const user = "root"
const pass = "pass"

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
		sortedHistory := sortHistory(history)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			pages.ErrorPage("Something went wrong.").Render(c, c.Writer)
			return
		}

		pages.UploadHistory(sortedHistory).Render(c, c.Writer)
	}
}

func sortHistory(history map[string][]string) []models.UploadHistory {
	var sortedHistory []models.UploadHistory

	var dates []string
	for date := range history {
		dates = append(dates, date)
	}

	sort.Slice(dates, func(i, j int) bool {
		date1, _ := time.Parse("2006-01-02", dates[i])
		date2, _ := time.Parse("2006-01-02", dates[j])
		return date2.Before(date1)
	})

	for _, date := range dates {
		sortedHistory = append(sortedHistory, models.UploadHistory{
			Date:   date,
			Images: history[date],
		})
	}

	return sortedHistory
}
