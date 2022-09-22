package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"os"
)

var (
	s3Source *s3.S3
)

func init() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("aws_region"))},
	)
	if err != nil {
		log.Fatalf("failed to create AWS session, %v", err)
	}
	s3Source = s3.New(sess)
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
	log.Println(message.Body)
	return nil
}
