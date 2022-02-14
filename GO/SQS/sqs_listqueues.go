package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

//// https://github.com/awsdocs/aws-doc-sdk-examples/tree/main/go/example_code/sqs
// GetQueues returns a list of queue names
func GetQueues(sess *session.Session) (*sqs.ListQueuesOutput, error) {
	// Create an SQS service client
	svc := sqs.New(sess)

	result, err := svc.ListQueues(nil)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func main() {
	// Create a session that gets credential values from ~/.aws/credentials
	// and the default region from ~/.aws/config
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	result, err := GetQueues(sess)
	if err != nil {
		fmt.Println("Got an error retrieving queue URLs:")
		fmt.Println(err)
		return
	}
	log.Println("I am here")
	for _, url := range result.QueueUrls {
		log.Println(*url)
	}
}
