# Cloud Vision API with Golang

## Prerequisite

You need to export a service account json file into `GOOGLE_APPLICATION_CREDENTIALS` env variable.

```
$ export GOOGLE_APPLICATION_CREDENTIALS=/path/to/service_account.json
```

Abort this execution file if you don't set the variable or unable to find the file.


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

## Example

### input

![lenna.jpg](https://raw.githubusercontent.com/kaneshin/go-cloud-vision/master/lenna.jpg)

### output

```
[
  {
    "angerLikelihood": "VERY_UNLIKELY",
    "blurredLikelihood": "VERY_UNLIKELY",
    "boundingPoly": {
      "vertices": [
        {
          "x": 143,
          "y": 43
        },
        {
          "x": 245,
          "y": 43
        },
        {
          "x": 245,
          "y": 163
        },
        {
          "x": 143,
          "y": 163
        }
      ]
    },
    "detectionConfidence": 0.99805844,
    "fdBoundingPoly": {
      "vertices": [
        {
          "x": 172,
          "y": 82
        },
        {
          "x": 241,
          "y": 82
        },
        {
          "x": 241,
          "y": 151
        },
        {
          "x": 172,
          "y": 151
        }
      ]
    },
    "headwearLikelihood": "UNLIKELY",
    "joyLikelihood": "VERY_UNLIKELY",
    "landmarkingConfidence": 0.5350582,
    "landmarks": [
      {
        "position": {
          "x": 197.90556,
          "y": 102.932,
          "z": 0.00083794753
        },
        "type": "LEFT_EYE"
      },
      {
        "position": {
          "x": 223.43489,
          "y": 102.72927,
          "z": 17.352478
        },
        "type": "RIGHT_EYE"
      },
      ...
      {
        "position": {
          "x": 167.89359,
          "y": 133.45166,
          "z": 18.304451
        },
        "type": "CHIN_LEFT_GONION"
      },
      {
        "position": {
          "x": 220.39294,
          "y": 132.60187,
          "z": 54.00494
        },
        "type": "CHIN_RIGHT_GONION"
      }
    ],
    "panAngle": 34.809193,
    "rollAngle": 5.9516206,
    "sorrowLikelihood": "VERY_UNLIKELY",
    "surpriseLikelihood": "VERY_UNLIKELY",
    "tiltAngle": -10.011973,
    "underExposedLikelihood": "VERY_UNLIKELY"
  }
]
```


## License

[The MIT License (MIT)](http://kaneshin.mit-license.org/)


## Author

Shintaro Kaneko <kaneshin0120@gmail.com>
