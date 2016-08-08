// MS Cognitive Services - Computer Vision API
// OCR service
package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"encoding/json"

	"github.com/BurntSushi/toml"
	"github.com/mondora/cvision-go/cvision"
	"gopkg.in/resty.v0"
	//"./cvision"
)

// FlagParams cmdline params
type FlagParams struct {
	cmd               string
	url               string
	filePath          string
	language          string
	detectOrientation bool
	jsonPrettyPrint   bool
	verbose           bool
	height            int
	width             int
}

// CvOcrConfig toml config
type CvOcrConfig struct {
	CvAPIKey1 string
	CvAPIKey2 string
}

// global vars
var config CvOcrConfig
var params FlagParams

// json pretty print format
func prettyprint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}

func readConfig() error {
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func init() {
	if err := readConfig(); err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	flag.StringVar(&params.cmd, "c", "tag", "command ocr|tag|analyze|describe|domain|recognize|thumbnail")
	flag.StringVar(&params.url, "url", "", "url")
	flag.StringVar(&params.url, "u", "", "url")
	flag.StringVar(&params.filePath, "file", "", "image path")
	flag.StringVar(&params.filePath, "f", "", "image path")
	flag.StringVar(&params.language, "l", "it", "language")
	flag.BoolVar(&params.jsonPrettyPrint, "pp", false, "json pretty print")
	flag.BoolVar(&params.detectOrientation, "d", true, "detect orientation")
	flag.BoolVar(&params.verbose, "v", false, "verbose")
	flag.IntVar(&params.height, "height", 0, "height")
	flag.IntVar(&params.width, "width", 0, "width")
	flag.Parse()
}

func main() {
	if params.verbose {
		log.Printf("Microsoft Cognitive Services - Computer Vision API %s", cvision.Version)
		log.Printf("> POST image: %s\n", params.filePath)
	}

	var resp *resty.Response
	var err error
	switch params.cmd {
	case "ocr": // OCR
		cvOCR := cvision.NewOCRClient(config.CvAPIKey1, params.language, params.detectOrientation)
		if params.url != "" {
			resp, err = cvOCR.GetOcrInfo(params.url, true, params.verbose)
		} else {
			resp, err = cvOCR.GetOcrInfo(params.filePath, false, params.verbose)
		}
	case "tag": // TAG
		cvTAG := cvision.NewTAGClient(config.CvAPIKey1)
		if params.url != "" {
			resp, err = cvTAG.GetTagInfo(params.url, true, params.verbose)
		} else {
			resp, err = cvTAG.GetTagInfo(params.filePath, false, params.verbose)
		}
	case "analyze": // Analyze Image (include TAG)
		cvAnalyze := cvision.NewAnalyzeClient(config.CvAPIKey1)
		if params.url != "" {
			resp, err = cvAnalyze.GetAnalyzeInfo(params.url, true, params.verbose)
		} else {
			resp, err = cvAnalyze.GetAnalyzeInfo(params.filePath, false, params.verbose)
		}
	case "describe": // Describe Image
		cvDescribe := cvision.NewDescribeClient(config.CvAPIKey1, nil)
		if params.url != "" {
			resp, err = cvDescribe.DescribeImage(params.url, true, params.verbose)
		} else {
			resp, err = cvDescribe.DescribeImage(params.filePath, false, params.verbose)
		}
	case "domain": // List Domain
		cvListDomain := cvision.NewListDomainClient(config.CvAPIKey1)
		resp, err = cvListDomain.GetListDomain(params.verbose)
		if params.verbose && err == nil {
			list, _ := cvision.UnmarshalListModel(resp.Body())
			l := len(list.Models)
			for i := 0; i < l; i++ {
				fmt.Printf("model found: %s\n", list.Models[i].Name)
			}
		}
	case "recognize": // Recognize Domain Specific Content
		model := "celebrities"
		cvRecognizeDomain := cvision.NewRecognizeDomainClient(config.CvAPIKey1, model)
		if params.url != "" {
			resp, err = cvRecognizeDomain.RecognizeDomain(params.url, true, params.verbose)
		} else {
			resp, err = cvRecognizeDomain.RecognizeDomain(params.filePath, false, params.verbose)
		}
	case "thumbnail": // Get Thumbnail
		smartCrop := true
		cvThumbnail := cvision.NewThumbnailClient(config.CvAPIKey1, params.width, params.height, &smartCrop)
		if params.url != "" {
			resp, err = cvThumbnail.GetThumbnail(params.url, true, params.verbose)
		} else {
			resp, err = cvThumbnail.GetThumbnail(params.filePath, false, params.verbose)
		}
	}
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode() == 200 {
		contentType := resp.Header().Get("Content-Type")
		if params.verbose {
			log.Printf("> resp code: %s Content-Type: %s\n", resp.Status(), contentType)
		}
		if strings.HasPrefix(contentType, "application/json") {
			if params.jsonPrettyPrint {
				b, _ := prettyprint(resp.Body())
				fmt.Printf("%s\n", b)
			} else {
				fmt.Printf("%s\n", resp.Body())
			}
		} else {
			fmt.Printf("%s", resp.Body())
		}
	} else {
		log.Printf("> resp code: %s\n", resp.Status())
		log.Printf("%s", resp)
	}
}
