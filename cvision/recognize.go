// Package cvision - MS Cognitive Services - Computer Vision API
// Recognize Domain Specific Content service
// https://dev.projectoxford.ai/docs/services/56f91f2d778daf23d8ec6739/operations/56f91f2e778daf14a499e200
package cvision

import (
	"fmt"
	"log"

	"gopkg.in/resty.v0"
)

const (
	// RecognizeDomainAPIURLBase Recognize Domain Specific Content Api url base
	// https://api.projectoxford.ai/vision/v1.0/models/{model}/analyze
	RecognizeDomainAPIURLBase = APIHost + "/vision/" + Version + "/models"
)

// RecognizeDomainClient struct
type RecognizeDomainClient struct {
	APIToken *string
	Model    string
}

// NewRecognizeDomainClient with parameters
func NewRecognizeDomainClient(token string, model string) *RecognizeDomainClient {
	return &RecognizeDomainClient{
		APIToken: &token,
		Model:    model,
	}
}

// RecognizeDomain recognizes content within an image by applying a domain-specific model
func (c *RecognizeDomainClient) RecognizeDomain(resourceName string, isURL bool, verbose bool) (*resty.Response, error) {
	if verbose {
		log.Printf("> RecognizeDomain: %s\n", resourceName)
	}
	apiURL := fmt.Sprintf("%s/%s/analyze", RecognizeDomainAPIURLBase, c.Model)
	req := resty.R().
		SetHeader("Ocp-Apim-Subscription-Key", *c.APIToken).
		SetHeader("User-Agent", UserAgent)
	if isURL {
		json, contentType := GetJSONURLPayload(resourceName)
		return req.
			SetHeader("Content-Type", contentType).
			SetBody(json).
			Post(apiURL)
	}
	return req.
		SetFile("recognize_img.jpeg", resourceName).
		Post(apiURL)
}
