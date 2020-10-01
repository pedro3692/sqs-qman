package qman

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func GetQueueURL(svc *sqs.SQS, queueName *string) (*sqs.GetQueueUrlOutput, error) {
	result, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: queueName,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetQueues(svc *sqs.SQS) (*sqs.ListQueuesOutput, error) {
	result, err := svc.ListQueues(nil)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func CreateQueue(svc *sqs.SQS, queue *string) (*sqs.CreateQueueOutput, error) {
	result, err := svc.CreateQueue(&sqs.CreateQueueInput{
		QueueName: queue,
		Attributes: map[string]*string{
			"DelaySeconds":           aws.String("60"),
			"MessageRetentionPeriod": aws.String("86400"),
		},
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func DeleteQueue(svc *sqs.SQS, queueURL *string) error {
	_, err := svc.DeleteQueue(&sqs.DeleteQueueInput{
		QueueUrl: queueURL,
	})

	if err != nil {
		return err
	}

	return nil
}
