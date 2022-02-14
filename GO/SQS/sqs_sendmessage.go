package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// Usage:
// go run sqs_sendmessage.go
func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess)

	// URL to our queue
	qURL := "https://sqs.us-west-2.amazonaws.com/787173426179/ANOR_LONDO"

	result, err := svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"phone_number": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String("45534545"),
			},
			"call_data": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String("IDDD"),
			},
			"agent_name": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String("Mathews"),
			},
			"costumer_id": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String("007"),
			},
			"product_id": &sqs.MessageAttributeValue{
				DataType:    aws.String("Number"),
				StringValue: aws.String("43321"),
			},
		},
		MessageBody: aws.String("event data"),
		QueueUrl:    &qURL,
	})

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Success", *result.MessageId)
}
