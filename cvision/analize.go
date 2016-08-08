// Package cvision - MS Cognitive Services - Computer Vision API
// Analyze Image service
// https://dev.projectoxford.ai/docs/services/56f91f2d778daf23d8ec6739/operations/56f91f2e778daf14a499e1fa
package cvision

import (
	"log"

	"gopkg.in/resty.v0"
)

const (
	// AnalyzeAPIURLBase Analyze Image Api url base
	AnalyzeAPIURLBase = APIHost + "/vision/" + Version + "/analyze"
)

// AnalyzeClient struct
type AnalyzeClient struct {
	APIToken *string
}

// NewAnalyzeClient with parameters
func NewAnalyzeClient(token string) *AnalyzeClient {
	return &AnalyzeClient{
		APIToken: &token,
	}
}

// GetAnalyzeInfo return json info related to image resourceName
func (c *AnalyzeClient) GetAnalyzeInfo(resourceName string, isURL bool, verbose bool) (*resty.Response, error) {
	if verbose {
		log.Printf("> GetAnalyzeInfo: %s\n", resourceName)
	}
	visualFeatures := "Categories,Tags,Description,Faces,ImageType,Color,Adult"
	details := "Celebrities"
	req := resty.R().
		SetQueryParams(map[string]string{
			"visualFeatures": visualFeatures,
			"details":        details,
		}).
		SetHeader("Ocp-Apim-Subscription-Key", *c.APIToken).
		SetHeader("User-Agent", UserAgent)
	if isURL {
		json, contentType := GetJSONURLPayload(resourceName)
		return req.
			SetHeader("Content-Type", contentType).
			SetBody(json).
			Post(AnalyzeAPIURLBase)
	}
	return req.
		SetFile("analyze_img.jpeg", resourceName).
		Post(AnalyzeAPIURLBase)
}
