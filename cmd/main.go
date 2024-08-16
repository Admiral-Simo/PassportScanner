package main

import (
	"PassportScanner/api/handler"
	"PassportScanner/scannersdk"
	"PassportScanner/views/pages"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	scanner := scannersdk.NewScannerSDK()
	contactStore := handler.NewContactStore()

	r.Static("/public", "./public")
	// pages
	r.GET("/", handler.PresentationPageHandler)
	r.GET("/main", handler.MainPageHandler)
	r.POST("/get-document-data", handler.DocumentDataPageHandler(scanner))
	r.GET("/authenticate", func(c *gin.Context) {
		pages.LoginPage(nil).Render(c, c.Writer)
	})
	r.POST("/upload-history", handler.UploadHistoryPageHandler(scanner))
	r.GET("/contact", handler.ContactPageHandler)
	r.POST("/contact/message/new", handler.PostDataContactPageHandler(contactStore))

	go cleanUpPublicFile()

	r.Run(":4000")
}

func cleanUpPublicFile() {
	for {
		time.Sleep(time.Hour)
		deleteUnnecessaryFiles()
	}
}

func deleteUnnecessaryFiles() {
	dir := "./public"

	_ = filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.Name() == "favicon.ico" {
			return nil
		}

		if !info.IsDir() {
			fileInfo, err := os.Stat(path)
			if err != nil {
				return err
			}

			now := time.Now()

			if now.Sub(fileInfo.ModTime()) > 5*time.Minute {
				err := os.Remove(path)
				if err != nil {
					fmt.Printf("Failed to delete file %s: %v\n", path, err)
				} else {
					fmt.Printf("Deleted file %s\n", path)
				}
			}
		}
		return nil
	})
}
