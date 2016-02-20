# Cloud Vision API with Golang

## Installation

Type the below command to install if you use this application on your device.

```shell
go get github.com/kaneshin/go-cloud-vision
```

Make sure that `go-cloud-vision` was installed correctly:

```shell
go-cloud-vision -h
```

### Dependencies

To run this sample, you need to install this packages;

- golang.org/x/net/context
- golang.org/x/oauth2/google
- google.golang.org/api/vision/...

This sample repository is contained glide.yaml to privide `glide` command. So you should install that packages with the `glide` command.

See https://github.com/Masterminds/glide


## Usage

```shell
go run main.go lenna.jpg
```

or if you already installed as a command.

```shell
go-cloud-vision lenna.jpg
```


## License

[The MIT License (MIT)](http://kaneshin.mit-license.org/)


## Author

Shintaro Kaneko <kaneshin0120@gmail.com>
