# cvision-go
Microsoft ComputerVision GO client

See https://www.microsoft.com/cognitive-services/en-us/computer-vision-api to get API KEYs

#### BUILD
```sh
go build -o cv
```

#### BUILD - Cross compile for Raspberry Pi
```sh
env GOOS=linux GOARCH=arm GOARM=6 go build -v -o cv-rpi
```

