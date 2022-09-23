package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
	"log"
	"os"
	"strings"
)

var (
	bucketName   string
	fileUploader *s3manager.Uploader
)

func init() {
	bucketName = os.Getenv("directoryobject_blobstore")
	fileUploader = initS3Uploader()
}

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, event events.SQSEvent) error {

	for _, item := range event.Records {
		err := processEvent(item)

		if err != nil {
			return err
		}
	}

	return nil
}

func processEvent(message events.SQSMessage) error {
	input := &s3manager.UploadInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(getFileKey(message)),
		Body:        strings.NewReader(message.Body),
		ContentType: aws.String("application/json"),
	}
	_, err := fileUploader.UploadWithContext(context.Background(), input)
	if err != nil {
		return err
	}

	return nil
}

func getFileKey(message events.SQSMessage) string {
	if message.MessageId != "" {
		return message.MessageId
	}

	return uuid.New().String()
}

func initS3Uploader() *s3manager.Uploader {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("aws_region"))},
	)
	if err != nil {
		log.Fatalf("failed to create AWS session, %v", err)
	}

	uploader := s3manager.NewUploader(sess)
	return uploader
}
