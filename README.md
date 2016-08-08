# cvision-go
Microsoft ComputerVision GO client

See https://www.microsoft.com/cognitive-services/en-us/computer-vision-api to get API KEYs

### How to install dependencies
```sh
$ go get github.com/mondora/cvision-go/cvision
$ go get github.com/BurntSushi/toml
$ go get gopkg.in/resty.v0
```
These packages will be installed on your [$GOPATH](https://golang.org/doc/install#install) location

### BUILD
```sh
$ go build -o cv
```

### BUILD - Cross compile for Raspberry Pi
```sh
$ env GOOS=linux GOARCH=arm GOARM=6 go build -v -o cv-rpi
```

### USAGE
First rename file `config.toml.demo` to `config.toml` and replace your API keys.
Basic usage: 
```sh
$ ./cv -u <URL_IMAGE> -c [ocr|tag|analyze|describe|recognize] [-pp] [-v]
$ ./cv -f <IMAGE_PATH> -c [ocr|tag|analyze|describe|recognize] [-pp] [-v]
$ ./cv -c domain [-pp] [-v]
```

### EXAMPLES
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

### FEATURES
- [x] ANALIZE
- [x] DESCRIBE
- [ ] THUMBNAIL
- [x] LIST DOMAIN
- [x] OCR
- [x] RECOGNIZE DOMAIN (recognize celebrities)
- [x] TAG
