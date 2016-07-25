// MS Cognitive Services - Computer Vision API
// OCR service
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"gopkg.in/resty.v0"

	"./cvision"
)

// FlagParams cmdline params
type FlagParams struct {
	cmd               string
	url               string
	filePath          string
	language          string
	detectOrientation bool
}

// CvOcrConfig toml config
type CvOcrConfig struct {
	CvAPIKey1 string
	CvAPIKey2 string
}

// global vars
var config CvOcrConfig
var params FlagParams

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
	flag.StringVar(&params.cmd, "c", "tag", "command")
	flag.StringVar(&params.url, "url", "", "url")
	flag.StringVar(&params.filePath, "file", "", "image path")
	flag.StringVar(&params.filePath, "f", "", "image path")
	flag.StringVar(&params.language, "l", "it", "language")
	flag.BoolVar(&params.detectOrientation, "d", true, "detect orientation")
	flag.Parse()
}

func main() {
	log.Printf("Microsoft Cognitive Services - Computer Vision API %s", cvision.Version)
	log.Printf("> POST image: %s\n", params.filePath)

	var resp *resty.Response
	var err error
	switch params.cmd {
	case "ocr": // OCR
		cvOCR := cvision.NewOCRClient(config.CvAPIKey1, params.language, params.detectOrientation)
		if params.url != "" {
			resp, err = cvOCR.GetOcrInfo(params.url, true)
		} else {
			resp, err = cvOCR.GetOcrInfo(params.filePath, false)
		}
	case "tag": // TAG
		cvTAG := cvision.NewTAGClient(config.CvAPIKey1)
		if params.url != "" {
			resp, err = cvTAG.GetTagInfo(params.url, true)
		} else {
			resp, err = cvTAG.GetTagInfo(params.filePath, false)
		}
	case "analyze": // Analyze Image
		cvAnalyze := cvision.NewAnalyzeClient(config.CvAPIKey1)
		if params.url != "" {
			resp, err = cvAnalyze.GetAnalyzeInfo(params.url, true)
		} else {
			resp, err = cvAnalyze.GetAnalyzeInfo(params.filePath, false)
		}
	}
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("> resp code: %s\n", resp.Status())
	if resp.StatusCode() == 200 {
		log.Printf("Body:\n%s\n", resp.Body())
	}
}
