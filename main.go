package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	// we expect exactly one CLI argument: the bucket name
	// if that is not provided, exit
	if len(os.Args) != 2 {
		log.Fatalf("Can't continue, no bucket name was provided!")
	}
	bucket := os.Args[1]
	now := time.Now()
	key := fmt.Sprintf("s3echoer-%v", now.Unix())

	userinput, err := userInput()
	if err != nil {
		log.Fatalf("Can't read from stdin: %v", err)
	}
	fmt.Printf("Uploading user input to S3 using %v/%v\n\n", bucket, key)

	err = uploadToS3(bucket, key, userinput)
	if err != nil {
		log.Fatalf("Can't upload to S3: %v", err)
	}
}

// userInput reads from stdin until it sees a CTRL+D.
func userInput() (string, error) {
	rawinput, err := io.ReadAll(os.Stdin)
	if err != nil {
		return "", err
	}
	return string(rawinput), nil
}

// uploadToS3 puts the payload into the S3 bucket using the key provided.
func uploadToS3(bucket, key, payload string) error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return err
	}
	client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(client)
	_, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   strings.NewReader(payload),
	})
	return err
}
