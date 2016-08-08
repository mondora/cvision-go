// Package cvision - MS Cognitive Services - Computer Vision API
// List Domain Specific Models service
// https://dev.projectoxford.ai/docs/services/56f91f2d778daf23d8ec6739/operations/56f91f2e778daf14a499e1fd
package cvision

import (
	"encoding/json"
	"log"

	"gopkg.in/resty.v0"
)

const (
	// ListDomainAPIURLBase List Domain Image Api url base
	ListDomainAPIURLBase = APIHost + "/vision/" + Version + "/models"
)

// ListDomainClient struct
type ListDomainClient struct {
	APIToken *string
}

/*
JSON example:
{"models":[{"name":"celebrities","categories":["people_"]}],"requestId":"584ac6eb-4521-4c42-b8d6-f5e7b75ef454"}
*/

// Model type
type Model struct {
	Name       string
	Categories []string
}

// ListModel type
type ListModel struct {
	Models    []Model
	RequestID string
}

// NewListDomainClient with parameters
func NewListDomainClient(token string) *ListDomainClient {
	return &ListDomainClient{
		APIToken: &token,
	}
}

// UnmarshalListModel from json to struct
func UnmarshalListModel(b []byte) (ListModel, error) {
	//b := []byte(jsonStr)
	var l ListModel
	err := json.Unmarshal(b, &l)
	return l, err
}

// GetListDomain return json List Domain Specific Models
func (c *ListDomainClient) GetListDomain(verbose bool) (*resty.Response, error) {
	if verbose {
		log.Printf("> GetListDomain\n")
	}
	req := resty.R().
		SetHeader("Ocp-Apim-Subscription-Key", *c.APIToken).
		SetHeader("User-Agent", UserAgent)
	return req.Get(ListDomainAPIURLBase)
}
