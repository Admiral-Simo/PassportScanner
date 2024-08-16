package handler

import (
	"PassportScanner/views/pages"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func ContactPageHandler(c *gin.Context) {
	pages.ContactPage(false).Render(c, c.Writer)
}

func PostDataContactPageHandler(store ContactStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract form data
		email := c.PostForm("email")
		phone := c.PostForm("phone")
		name := c.PostForm("name")
		companyName := c.PostForm("company_name")
		message := c.PostForm("message")

		// Format the data for storage
		data := fmt.Sprintf("Email: %s\nPhone: %s\nName: %s\nCompany Name: %s\nMessage: %s\n\n",
			email, phone, name, companyName, message)

		err := store.StoreDataInFile(data)
		if err != nil {
			c.Status(500)
			pages.ErrorPage("Unable to write data to file").Render(c, c.Writer)
			return
		}

		// Respond to the client
		pages.ContactPage(true).Render(c, c.Writer)
	}
}

type ContactStore struct {
	File *os.File
}

func NewContactStore() ContactStore {
	file, err := os.Create("contact_page_data.txt")
	if err != nil {
		log.Fatal(err)
	}
	return ContactStore{File: file}
}

func (c ContactStore) StoreDataInFile(data string) error {
	_, err := c.File.WriteString(data)
	if err != nil {
		return err
	}

	return nil
}

func (c ContactStore) ReadDataFromFile() (string, error) {
	var data strings.Builder
	scanner := bufio.NewScanner(c.File)
	for scanner.Scan() {
		data.WriteString(scanner.Text() + "\n")
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return data.String(), nil
}
