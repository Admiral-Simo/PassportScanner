package scannersdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

func (sdk *ScannerSDK) GetDocumentData(file *multipart.FileHeader) (*DocumentData, error) {
	// Open the file
	f, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Create a buffer to hold the multipart data
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	// Create a form file field and write the file to it
	fw, err := w.CreateFormFile("file", file.Filename)
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(fw, f); err != nil {
		return nil, err
	}

	// Close the writer to set the terminating boundary
	w.Close()

	// Send the POST request
	resp, err := http.Post(sdk.endpoint+"/get-document-data", w.FormDataContentType(), &b)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON response into the DocumentData struct
	var documentData DocumentData
	if err := json.Unmarshal(responseData, &documentData); err != nil {
		return nil, err
	}

	return &documentData, nil
}

// getUploadHistory makes an HTTP GET request to the endpoint and unmarshals the response
func (sdk *ScannerSDK) GetUploadHistory() (map[string][]string, error) {
	// Make the HTTP GET request
	resp, err := http.Get(sdk.endpoint + "/get-upload-history")
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	// Check for non-200 HTTP status codes
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	// Unmarshal the JSON response into a map
	var result map[string][]string
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON response: %v", err)
	}

	return result, nil
}
