package handler

import (
	"PassportScanner/views/pages"

	"github.com/gin-gonic/gin"
)

func PresentationPageHandler(c *gin.Context) {
	pages.PresentationPage().Render(c, c.Writer)
}
