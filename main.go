package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
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

	awsEndpoint := "http://localhost:4566"
	awsRegion := "us-east-2"

	if *msg == "" || *topicARN == "" {
		fmt.Println("You must supply a message and topic ARN")
		fmt.Println("-m MESSAGE -t TOPIC-ARN")
		return
	}

	customResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		if awsEndpoint != "" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           awsEndpoint,
				SigningRegion: awsRegion,
			}, nil
		}

		// returning EndpointNotFoundError will allow the service to fallback to its default resolution
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(awsRegion),
		config.WithEndpointResolver(customResolver),
	)
	if err != nil {
		log.Fatalf("Cannot load the AWS configs: %s", err)
	}

	//--------------------------------------------------------------------

	client := sns.NewFromConfig(cfg)

	result, err := client.ListSubscriptions(context.TODO(), &sns.ListSubscriptionsInput{})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("start")
	for _, t := range result.Subscriptions {
		fmt.Println(*t.TopicArn)
	}
	fmt.Println("end")
	//--------------------------------------------------------------------

	// //--------------------------------------------------------------------
	// client := sns.NewFromConfig(cfg)

	// result, err := client.Publish(context.TODO(), &sns.PublishInput{
	// 	Message:  msg,
	// 	TopicArn: topicARN,
	// })

	// if err != nil {
	// 	fmt.Println("Got an error publishing the message:")
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println("Message ID: " + *result.MessageId)
	// //--------------------------------------------------------------------

	// //--------------------------------------------------------------------

	// fmt.Println("List all buckets.")

	// // Create the resource client
	// client := s3.NewFromConfig(cfg, func(o *s3.Options) {
	// 	o.UsePathStyle = true
	// })

	// buckets, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	// if err != nil {
	// 	fmt.Printf("Couldn't list buckets: %v", err)
	// 	return
	// }

	// for _, bucket := range buckets.Buckets {
	// 	fmt.Printf("Found bucket: %s, created at: %s\n", *bucket.Name, *bucket.CreationDate)
	// }

	// //--------------------------------------------------------------------
}

// //===============================================================

// import (
// 	"context"
// 	"fmt"

// 	"github.com/aws/aws-sdk-go-v2/aws"
// 	"github.com/aws/aws-sdk-go-v2/service/sns"
// )

// type SNSPublishAPI interface {
// 	Publish(ctx context.Context,
// 		params *sns.PublishInput,
// 		optFns ...func(*sns.Options)) (*sns.PublishOutput, error)
// }

// func PublishMessage(c context.Context, api SNSPublishAPI, input *sns.PublishInput) (*sns.PublishOutput, error) {
// 	return api.Publish(c, input)
// }

// func main() {
// 	// msg := flag.String("m", "", "The message to send to the subscribed users of the topic")
// 	// topicARN := flag.String("t", "", "The ARN of the topic to which the user subscribes")

// 	msg := aws.String("test message")
// 	topicARN := aws.String("arn:aws:sns:us-west-2:000000000000:first-proj1")

// 	// flag.Parse()

// 	if *msg == "" || *topicARN == "" {
// 		fmt.Println("You must supply a message and topic ARN")
// 		fmt.Println("-m MESSAGE -t TOPIC-ARN")
// 		return
// 	}
// 	cfg := aws.Config{
// 		EndpointResolver: aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
// 			return aws.Endpoint{
// 				PartitionID:       "aws",
// 				URL:               "http://localhost:4566",
// 				SigningRegion:     "us-west-2",
// 				HostnameImmutable: true,
// 			}, nil
// 		}),
// 	}

// 	client := sns.NewFromConfig(cfg)
// 	input := &sns.PublishInput{
// 		Message:  msg,
// 		TopicArn: topicARN,
// 	}

// 	result, err := PublishMessage(context.TODO(), client, input)
// 	if err != nil {
// 		fmt.Println("Got an error publishing the message:")
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println("Message ID: " + *result.MessageId)
// }
