package handlers

import (
	"passportScanner/views/pages"

	"github.com/gin-gonic/gin"
)

func ContactPageHandler(c *gin.Context) {
	pages.ContactPage(false).Render(c, c.Writer)
}
