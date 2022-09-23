package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
	"github.com/jrolstad/ldap-api/internal/pkg/core"
	"github.com/jrolstad/ldap-api/internal/pkg/models"
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
	directoryObject := &models.DirectoryObject{}
	core.MapFromJson(message.Body, directoryObject)

	input := &s3manager.UploadInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(getFileKey(directoryObject)),
		Body:        strings.NewReader(message.Body),
		ContentType: aws.String("application/json"),
	}
	_, err := fileUploader.UploadWithContext(context.Background(), input)
	if err != nil {
		return err
	}

	return nil
}

func getFileKey(item *models.DirectoryObject) string {
	path := resolveItemPath(item)
	identifier := resolveItemId(item)

	return fmt.Sprintf("%v/%v", path, identifier)
}

func resolveItemPath(item *models.DirectoryObject) string {
	if item == nil || item.ObjectType == "" {
		return "unknown"
	}

	return strings.ToLower(item.ObjectType)
}

func resolveItemId(item *models.DirectoryObject) string {
	if item == nil || item.Id == "" {
		return uuid.New().String()
	}

	return strings.ToLower(item.Id)
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
