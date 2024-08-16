package main

import (
	"PassportScanner/api"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

func main() {
	webAPI := api.New()

	go cleanUpPublicFile()

	webAPI.StartServer()
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
