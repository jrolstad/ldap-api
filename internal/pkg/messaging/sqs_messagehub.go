package messaging

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/jrolstad/ldap-api/internal/pkg/core"
)

type SqsMessageHub struct {
	sqs *sqs.SQS
}

func (hub *SqsMessageHub) Send(toSend interface{}, target string) error {
	message, mapError := hub.mapToSqsSendMessage(hub.sqs, toSend, target)
	if mapError != nil {
		return mapError
	}

	_, sendError := hub.sqs.SendMessage(message)
	if sendError != nil {
		return sendError
	}

	return nil
}

func (hub *SqsMessageHub) mapToSqsSendMessage(sqsInstance *sqs.SQS, toMap interface{}, queueName string) (*sqs.SendMessageInput, error) {
	urlResult, queueUrlErr := sqsInstance.GetQueueUrl(&sqs.GetQueueUrlInput{QueueName: aws.String(queueName)})
	if queueUrlErr != nil {
		return nil, queueUrlErr
	}

	input := new(sqs.SendMessageInput)
	input.MessageBody = aws.String(core.MapToJson(toMap))
	input.QueueUrl = urlResult.QueueUrl

	return input, nil
}
