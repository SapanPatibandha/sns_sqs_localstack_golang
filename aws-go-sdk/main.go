package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func main() {

	msg := aws.String("test message")
	topicARN := aws.String("arn:aws:sns:us-west-2:000000000000:trial-proj1-sns")

	// Initialize a session
	sess, _ := session.NewSession(&aws.Config{
		Region:           aws.String("us-west-2"),
		Credentials:      credentials.NewStaticCredentials("test", "test", ""),
		S3ForcePathStyle: aws.Bool(true),
		Endpoint:         aws.String("http://localhost:4566"),
	})

	//--------------------------------------------------------------------
	client := sns.New(sess)

	input := &sns.PublishInput{
		Message:  msg,
		TopicArn: topicARN,
	}

	result, err := client.Publish(input)

	if err != nil {
		fmt.Println("Got an error publishing the message:")
		fmt.Println(err)
		return
	}

	fmt.Println("Message ID: " + *result.MessageId)
	//--------------------------------------------------------------------
}
