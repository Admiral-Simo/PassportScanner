package handlers

import (
	"bufio"
	"fmt"
	"os"
	"passportScanner/views/pages"
	"strings"

	"github.com/gin-gonic/gin"
)

func ContactPageHandler(c *gin.Context) {
	pages.ContactPage(false).Render(c, c.Writer)
}

func PostDataContactPageHandler(c *gin.Context) {
	// Extract form data
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	name := c.PostForm("name")
	companyName := c.PostForm("company_name")
	message := c.PostForm("message")

	// Format the data for storage
	data := fmt.Sprintf("Email: %s\nPhone: %s\nName: %s\nCompany Name: %s\nMessage: %s\n\n",
		email, phone, name, companyName, message)

	// Store data in a file
	fileName := "contact_page_data.txt"
	err := StoreDataInFile(fileName, data)
	if err != nil {
		c.Status(500)
		pages.ErrorPage("Unable to write data to file").Render(c, c.Writer)
		return
	}

	// Respond to the client
	pages.ContactPage(true).Render(c, c.Writer)
}

func StoreDataInFile(fileName, data string) error {
	// Open the file in append mode
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return err
	}

	fmt.Printf("Data successfully written to %s\n", fileName)
	return nil
}

func ReadDataFromFile(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var data strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data.WriteString(scanner.Text() + "\n")
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return data.String(), nil
}
