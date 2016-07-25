// Package cvision - MS Cognitive Services - Computer Vision API
// TAG service
// https://dev.projectoxford.ai/docs/services/56f91f2d778daf23d8ec6739/operations/56f91f2e778daf14a499e1ff
package cvision

import (
	"log"

	"gopkg.in/resty.v0"
)

const (
	// TagAPIURLBase TAG Api url base
	TagAPIURLBase = APIHost + "/vision/" + Version + "/tag"
)

// TAGClient struct
type TAGClient struct {
	APIToken    *string
	contentType *string
	Verbose     bool
}

// NewTAGClient with parameters
func NewTAGClient(token string) *TAGClient {
	return &TAGClient{
		APIToken:    &token,
		contentType: nil,
		Verbose:     true,
	}
}

// GetTagInfo return json info related to image resourceName
func (c *TAGClient) GetTagInfo(resourceName string, isURL bool) (*resty.Response, error) {
	if c.Verbose {
		log.Printf("> GetTagInfo: %s\n", resourceName)
	}
	req := resty.R().
		SetHeader("Ocp-Apim-Subscription-Key", *c.APIToken).
		SetHeader("User-Agent", UserAgent)

	if isURL {
		json, contentType := GetJSONURLPayload(resourceName)
		return req.
			SetHeader("Content-Type", contentType).
			SetBody(json).
			Post(TagAPIURLBase)
	}
	return req.
		SetFile("tag_img.jpeg", resourceName).
		Post(TagAPIURLBase)
}
