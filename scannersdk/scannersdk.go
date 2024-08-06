package scannersdk

type ScannerSDK struct {
	endpoint string
}

func NewScannerSDK() ScannerSDK {
	return ScannerSDK{
		endpoint: "http://31.220.89.29:8091",
	}
}

// DocumentData holds the parsed document data
type DocumentData struct {
	DocumentNumber string `json:"documentNumber"`
	DocumentType   string `json:"documentType"`
	CountryCode    string `json:"countryCode"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Sex            string `json:"sex"`
	BirthDate      string `json:"birthDate"`
	ExpireDate     string `json:"expireDate"`
}
