# cvision-go
Microsoft ComputerVision GO client

[![GoDoc](https://godoc.org/github.com/mondora/cvision-go?status.svg)](https://godoc.org/github.com/mondora/cvision-go) [![Build Status](https://travis-ci.org/mondora/cvision-go.svg?branch=master)](https://travis-ci.org/mondora/cvision-go) [![Go Report Card](https://goreportcard.com/badge/github.com/mondora/cvision-go)](https://goreportcard.com/report/github.com/mondora/cvision-go)

## Overview
MS Cognitive Services - Computer Vision API (including OCR service)

This library includes a simple but complete command-line client `cv`

See https://www.microsoft.com/cognitive-services/en-us/computer-vision-api to get API KEYs

## Features
- [x] ANALIZE
- [x] DESCRIBE
- [x] THUMBNAIL
- [x] LIST DOMAIN
- [x] OCR
- [x] RECOGNIZE DOMAIN (recognize celebrities)
- [x] TAG

## Install
```sh
$ go get -u github.com/mondora/cvision-go/cvision
```
This package will be installed on your [$GOPATH](https://golang.org/doc/install#install) location

## Build
```sh
$ go build -o cv
```

## Build - Cross compile for Raspberry Pi
```sh
$ env GOOS=linux GOARCH=arm GOARM=6 go build -v -o cv-rpi
```

## Usage
First rename file `config.toml.demo` to `config.toml` and replace your API keys.
Basic usage: 
```sh
$ ./cv -u <URL_IMAGE> -c [ocr|tag|analyze|describe|recognize] [-pp] [-v]
$ ./cv -f <IMAGE_PATH> -c [ocr|tag|analyze|describe|recognize] [-pp] [-v]
$ ./cv -c domain [-pp] [-v]
$ ./cv [-u <URL_IMAGE>|-f <IMAGE_PATH>] -c thumbnail [-whidth <###>] [-height <###>] [-pp] [-v]
```

## Examples
```sh
$ ./cv -u http://www.onegossip.it/wp-content/uploads/2013/12/Michael_Schumacher.jpg -c recognize -pp
{
  "requestId": "273ed75e-d19c-4660-8f16-1f319d7e41ff",
  "metadata": {
    "width": 300,
    "height": 300,
    "format": "Jpeg"
  },
  "result": {
    "celebrities": [
      {
        "name": "Michael Schumacher",
        "faceRectangle": {
          "left": 96,
          "top": 61,
          "width": 99,
          "height": 99
        },
        "confidence": 0.9999989
      }
    ]
  }
}
```

```
$ ./cv -u "http://img2.tgcom24.mediaset.it/binary/articolo/instagram/94.\$plit/C_2_articolo_3024767_upiImagepp.jpg" -pp -v
2016/08/08 11:25:27 Microsoft Cognitive Services - Computer Vision API v1.0
2016/08/08 11:25:27 > POST image: 
2016/08/08 11:25:27 > GetTagInfo: http://img2.tgcom24.mediaset.it/binary/articolo/instagram/94.$plit/C_2_articolo_3024767_upiImagepp.jpg
2016/08/08 11:25:30 > resp code: 200 OK
{
  "tags": [
    {
      "name": "person",
      "confidence": 0.99995660781860352
    },
    {
      "name": "woman",
      "confidence": 0.99320459365844727
    },
    {
      "name": "smiling",
      "confidence": 0.83515453338623047
    },
    {
      "name": "lady",
      "confidence": 0.68646234273910522
    },
    {
      "name": "beautiful",
      "confidence": 0.37140417098999023
    },
    {
      "name": "pretty",
      "confidence": 0.36113372445106506
    }
  ],
  "requestId": "b611ae81-be33-426f-b58b-6d9d31500b84",
  "metadata": {
    "width": 597,
    "height": 336,
    "format": "Jpeg"
  }
}

$ ./cv -u "http://img2.tgcom24.mediaset.it/binary/articolo/instagram/94.\$plit/C_2_articolo_3024767_upiImagepp.jpg" -pp -c recognize
{
  "requestId": "71eff647-dd7c-417f-b015-b81ed3f45654",
  "metadata": {
    "width": 597,
    "height": 336,
    "format": "Jpeg"
  },
  "result": {
    "celebrities": [
      {
        "name": "Miley Cyrus",
        "faceRectangle": {
          "left": 201,
          "top": 86,
          "width": 186,
          "height": 186
        },
        "confidence": 0.9590725
      }
    ]
  }
}
```

```
$ ./cv -u http://hq-wall.net/i/med_thumb/05/65/Natalie_Imbruglia_6640d0d2674d4e2cca45330ad4f9ee37.jpg -c thumbnail -height 300 -pp -v > nat.jpg
2016/08/08 15:55:33 Microsoft Cognitive Services - Computer Vision API v1.0
2016/08/08 15:55:33 > POST image: 
2016/08/08 15:55:33 > GetThumbnail: http://hq-wall.net/i/med_thumb/05/65/Natalie_Imbruglia_6640d0d2674d4e2cca45330ad4f9ee37.jpg
2016/08/08 15:55:33   width: 485 height: 300
2016/08/08 15:55:54 > resp code: 200 OK Content-Type: image/jpeg
```

## Author

[Marco Rozzati](https://github.com/marco-rozz)

## License

MIT.
