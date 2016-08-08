// Package cvision - MS Cognitive Services - Computer Vision API
// Get Thumbnail service
// https://dev.projectoxford.ai/docs/services/56f91f2d778daf23d8ec6739/operations/56f91f2e778daf14a499e1fb
package cvision

import (
	"log"
	"strconv"

	"gopkg.in/resty.v0"
)

const (
	// ThumbnailAPIURLBase TAG Api url base
	ThumbnailAPIURLBase = APIHost + "/vision/" + Version + "/generateThumbnail"
)

// ThumbnailClient struct
type ThumbnailClient struct {
	APIToken      *string
	Width         int
	Height        int
	SmartCropping *bool
}

// NewThumbnailClient with parameters
func NewThumbnailClient(token string, w int, h int, sc *bool) *ThumbnailClient {
	return &ThumbnailClient{
		APIToken:      &token,
		Width:         w,
		Height:        h,
		SmartCropping: sc,
	}
}

func adjustSize(width int, height int) (int, int) {
	zeroCount := 0
	if height == 0 {
		zeroCount = zeroCount + 1
	}
	if width == 0 {
		zeroCount = zeroCount + 1
	}
	if zeroCount == 1 {
		goldenRatio := 1.618
		if height == 0 {
			height = int(float64(width) * goldenRatio)
		} else {
			width = int(float64(height) * goldenRatio)
		}
	}
	return width, height
}

// GetThumbnail return [smart] thumbnail related to original image resourceName
func (c *ThumbnailClient) GetThumbnail(resourceName string, isURL bool, verbose bool) (*resty.Response, error) {
	w, h := adjustSize(c.Width, c.Height)
	if verbose {
		log.Printf("> GetThumbnail: %s\n", resourceName)
		log.Printf("  width: %d height: %d", w, h)
	}
	req := resty.R().
		SetHeader("Ocp-Apim-Subscription-Key", *c.APIToken).
		SetHeader("User-Agent", UserAgent).
		SetQueryParams(map[string]string{
			"width":  strconv.Itoa(w),
			"height": strconv.Itoa(h),
		})
	if c.SmartCropping != nil {
		req.SetQueryParams(map[string]string{
			"smartCropping": strconv.FormatBool(*c.SmartCropping),
		})
	}
	if isURL {
		json, contentType := GetJSONURLPayload(resourceName)
		return req.
			SetHeader("Content-Type", contentType).
			SetBody(json).
			Post(ThumbnailAPIURLBase)
	}
	return req.
		SetFile("thumbnail_img.jpeg", resourceName).
		Post(ThumbnailAPIURLBase)
}
