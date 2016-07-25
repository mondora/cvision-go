// Package cvision - MS Cognitive Services - Computer Vision API
package cvision

import "fmt"

const (
	// Version # of cvision
	Version = "v1.0"
	// UserAgent of client
	UserAgent = "CVision GO client " + Version
	// APIHost of services
	APIHost = "https://api.projectoxford.ai"
)

// GetJSONURLPayload return json payload and content type
func GetJSONURLPayload(url string) (string, string) {
	json := fmt.Sprintf("{\"url\":\"%s\"}", url)
	return json, "application/json"
}
