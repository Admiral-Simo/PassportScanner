package api

import (
	"PassportScanner/api/handler"
	"PassportScanner/views/pages"

	"github.com/gin-gonic/gin"
)

func (a *API) registerRoutes() {
	// pages
	a.app.GET("/", handler.PresentationPageHandler)
	a.app.GET("/main", handler.MainPageHandler)
	a.app.POST("/get-document-data", handler.DocumentDataPageHandler(a.scanner))
	a.app.GET("/authenticate", func(c *gin.Context) {
		pages.LoginPage(nil).Render(c, c.Writer)
	})
	a.app.POST("/upload-history", handler.UploadHistoryPageHandler(a.scanner))
	a.app.GET("/contact", handler.ContactPageHandler)
	a.app.POST("/contact/message/new", handler.PostDataContactPageHandler(a.contactStore))
}
