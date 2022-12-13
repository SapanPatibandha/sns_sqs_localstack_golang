package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type SNSPublishAPI interface {
	Publish(ctx context.Context,
		params *sns.PublishInput,
		optFns ...func(*sns.Options)) (*sns.PublishOutput, error)
}

func PublishMessage(c context.Context, api SNSPublishAPI, input *sns.PublishInput) (*sns.PublishOutput, error) {
	return api.Publish(c, input)
}

func main() {
	// msg := flag.String("m", "", "The message to send to the subscribed users of the topic")
	// topicARN := flag.String("t", "", "The ARN of the topic to which the user subscribes")

	msg := aws.String("test message")
	topicARN := aws.String("arn:aws:sns:us-west-2:000000000000:trial-proj1-sns")

	// flag.Parse()

	if *msg == "" || *topicARN == "" {
		fmt.Println("You must supply a message and topic ARN")
		fmt.Println("-m MESSAGE -t TOPIC-ARN")
		return
	}
	cfg := aws.Config{
		EndpointResolver: aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           "http://localhost:4566",
				SigningRegion: "us-west-2",
				// HostnameImmutable: false,
			}, nil
		}),
	}

	client := sns.NewFromConfig(cfg)
	// input := &sns.PublishInput{
	// 	Message:          msg,
	// 	TopicArn:         topicARN,
	// 	MessageStructure: aws.String("json"),
	// }

	// // List SNS
	// result, err := client.Subscribe(context.TODO(), &sns.SubscribeInput{
	// 	Endpoint:              aws.String("sapan.patibandha@contis.com"),
	// 	Protocol:              aws.String("email"),
	// 	ReturnSubscriptionArn: *aws.Bool(true), // Return the ARN, even if user has yet to confirm
	// 	TopicArn:              topicARN,
	// })

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	os.Exit(1)
	// }

	// fmt.Println(*result.SubscriptionArn)

	// result, err := PublishMessage(context.TODO(), client, input)
	result, err := client.Publish(context.TODO(), &sns.PublishInput{
		Message:          msg,
		TopicArn:         topicARN,
		MessageStructure: aws.String("json"),
	})

	if err != nil {
		fmt.Println("Got an error publishing the message:")
		fmt.Println(err)
		return
	}

	fmt.Println("Message ID: " + *result.MessageId)
}
