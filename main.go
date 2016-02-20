package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	vision "google.golang.org/api/vision/v1"
)

const (
	SUCCESS = 0
	FAILURE = 1
)

func newVisionService() (*vision.Service, error) {
	ctx := context.Background()

	fn := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if len(fn) == 0 {
		return nil, fmt.Errorf("Unable to get env variable: GOOGLE_APPLICATION_CREDENTIALS")
	}

	b, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, fmt.Errorf("Unable to read client secret file: %v", err)
	}

	config, err := google.JWTConfigFromJSON(b, vision.CloudPlatformScope)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse client secret file to config: %v", err)
	}
	client := config.Client(ctx)

	srv, err := vision.New(client)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve vision Client %v", err)
	}
	return srv, nil
}

func annotateImageRequest(path string) (*vision.AnnotateImageRequest, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Unable to read an image by file path: %v", err)
	}

	req := &vision.AnnotateImageRequest{
		// Apply image which is encoded by base64.
		Image: &vision.Image{
			Content: base64.StdEncoding.EncodeToString(b),
		},
		// Apply features to indicate what type of image detection.
		Features: []*vision.Feature{
			{
				MaxResults: 10,
				Type:       "FACE_DETECTION",
			},
		},
	}
	return req, nil
}

func run() int {
	// Parse arguments to run this function.
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		log.Printf("Argument is required.")
		return FAILURE
	}

	// Initialize vision service by a credentials json.
	srv, err := newVisionService()
	if err != nil {
		log.Printf("Unable to retrieve vision service: %v\n", err)
		return FAILURE
	}

	// Create request by given argument (last one).
	req, err := annotateImageRequest(args[len(args)-1])
	if err != nil {
		log.Printf("Unable to retrieve image request: %v\n", err)
		return FAILURE
	}

	// To call multiple image annotation requests.
	batch := &vision.BatchAnnotateImagesRequest{
		Requests: []*vision.AnnotateImageRequest{req},
	}

	// Execute the "vision.images.annotate".
	res, err := srv.Images.Annotate(batch).Do()
	if err != nil {
		log.Printf("Unable to execute images annotate requests: %v\n", err)
		return FAILURE
	}

	// Marshal annotations from responses
	body, err := json.MarshalIndent(res.Responses, "", "  ")
	if err != nil {
		log.Printf("Unable to marshal the response: %v\n", err)
		return FAILURE
	}
	fmt.Println(string(body))

	return SUCCESS
}

func main() {
	os.Exit(run())
}
