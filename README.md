# cvision-go
Microsoft ComputerVision GO client

See https://www.microsoft.com/cognitive-services/en-us/computer-vision-api to get API KEYs

### How to install dependencies
```sh
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
```sh
$ ./cv -u <URL_IMAGE> -c [analyze|ocr|tag] [-pp]
$ ./cv -f <IMAGE_PATH> -c [analyze|ocr|tag] [-pp]
```

### FEATURES
- [x] ANALIZE
- [ ] DESCRIBE
- [ ] THUMBNAIL
- [ ] LIST DOMAIN
- [x] OCR
- [ ] RECOGNIZE DOMAIN
- [x] TAG
