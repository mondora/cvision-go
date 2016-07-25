// Package cvision - MS Cognitive Services - Computer Vision API
// OCR service
// https://dev.projectoxford.ai/docs/services/56f91f2d778daf23d8ec6739/operations/56f91f2e778daf14a499e1fc
package cvision

import (
	"fmt"
	"log"

	"gopkg.in/resty.v0"
)

const (
	// OcrAPIURLBase OCR Api url base
	OcrAPIURLBase = APIHost + "/vision/" + Version + "/ocr"
)

// OCRClient struct
type OCRClient struct {
	APIToken          *string
	language          *string
	detectOrientation bool
	Verbose           bool
}

// NewOCRClient with parameters
func NewOCRClient(token string, language string, detectOrientation bool) *OCRClient {
	return &OCRClient{
		APIToken:          &token,
		language:          &language,
		detectOrientation: detectOrientation,
		Verbose:           true,
	}
}

// GetOcrInfo return json info related to image resourceName
func (c *OCRClient) GetOcrInfo(resourceName string, isURL bool) (*resty.Response, error) {
	lang := *c.language
	detOrientation := fmt.Sprintf("%t", c.detectOrientation)
	if c.Verbose {
		log.Printf("> GetOcrInfo: %s\n", resourceName)
	}
	req := resty.R().
		SetQueryParams(map[string]string{
			"language":          lang,
			"detectOrientation": detOrientation,
		}).
		SetHeader("Ocp-Apim-Subscription-Key", *c.APIToken).
		SetHeader("User-Agent", UserAgent)
	if isURL {
		json, contentType := GetJSONURLPayload(resourceName)
		return req.
			SetHeader("Content-Type", contentType).
			SetBody(json).
			Post(OcrAPIURLBase)
	}
	return req.
		SetFile("ocr_img.jpeg", resourceName).
		Post(OcrAPIURLBase)
}
