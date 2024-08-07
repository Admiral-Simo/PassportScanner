package handlers

import (
	"PassportScanner/views/pages"

	"github.com/gin-gonic/gin"
)

func MainPageHandler(c *gin.Context) {
	pages.Main().Render(c, c.Writer)
}
