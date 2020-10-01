package main

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	qman "github.com/pedro3692/sqs-qman/lib"
)

func main() {

	queue := flag.String("q", "", "The name of the queue to create")
	list := flag.Bool("l", false, "List queues")
	delete := flag.String("d", "", "The name of the queue to delete")
	endpoint := flag.String("e", "http://127.0.0.1:9324", "SQS endpoint")

	flag.Parse()

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("elasticmq"),
		Credentials: credentials.NewStaticCredentials("x", "x", "x"),
	})

	svc := sqs.New(sess, &aws.Config{Endpoint: aws.String(*endpoint)})

	if err != nil {
		fmt.Println("Got an error creating the session:")
		fmt.Println(err)
		return
	}

	if !(*queue == "") {
		result, err := qman.CreateQueue(svc, queue)
		if err != nil {
			fmt.Println("Got an error creating the queue:")
			fmt.Println(err)
			return
		}

		fmt.Println("URL: " + *result.QueueUrl)
	}

	if *list {
		queueList, err := qman.GetQueues(svc)
		if err != nil {
			fmt.Println("Got an error retrieving queue URLs:")
			fmt.Println(err)
			return
		}

		fmt.Println("Queue List: ")

		for i, url := range queueList.QueueUrls {
			fmt.Printf("%d: %s\n", i, *url)
		}
	}

	if *delete != "" {
		result, err := qman.GetQueueURL(svc, delete)
		if err != nil {
			fmt.Println("Got an error getting the queue URL:")
			fmt.Println(err)
			return
		}

		queueURL := result.QueueUrl

		err = qman.DeleteQueue(svc, queueURL)
		if err != nil {
			fmt.Println("Got an error deleting the queue:")
			fmt.Println(err)
			return
		}

		fmt.Println("Deleted queue with URL " + *queueURL)
	}

}
