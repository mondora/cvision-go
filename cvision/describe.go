// Package cvision - MS Cognitive Services - Computer Vision API
// Describe service
// https://dev.projectoxford.ai/docs/services/56f91f2d778daf23d8ec6739/operations/56f91f2e778daf14a499e1fe
package cvision

import (
	"log"
	"strconv"

	"gopkg.in/resty.v0"
)

const (
	// DescribeAPIURLBase Describe Api url base
	DescribeAPIURLBase = APIHost + "/vision/" + Version + "/describe"
)

// DescribeClient struct
type DescribeClient struct {
	APIToken      *string
	MaxCandidates *int
}

// NewDescribeClient with parameters
func NewDescribeClient(token string, maxCandidates *int) *DescribeClient {
	return &DescribeClient{
		APIToken:      &token,
		MaxCandidates: maxCandidates,
	}
}

// DescribeImage return description (json) related to image resourceName
func (c *DescribeClient) DescribeImage(resourceName string, isURL bool, verbose bool) (*resty.Response, error) {
	if verbose {
		log.Printf("> DescribeImage: %s\n", resourceName)
	}
	req := resty.R().
		SetHeader("Ocp-Apim-Subscription-Key", *c.APIToken).
		SetHeader("User-Agent", UserAgent)
	if c.MaxCandidates != nil {
		req.SetQueryParams(map[string]string{
			"maxCandidates": strconv.Itoa(*c.MaxCandidates),
		})
	}
	if isURL {
		json, contentType := GetJSONURLPayload(resourceName)
		return req.
			SetHeader("Content-Type", contentType).
			SetBody(json).
			Post(DescribeAPIURLBase)
	}
	return req.
		SetFile("describe_img.jpeg", resourceName).
		Post(DescribeAPIURLBase)
}
