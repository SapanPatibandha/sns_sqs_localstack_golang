# AWS SNS and SQS using localstack with golang

Command to run localstack on Docker

```cli
docker run --rm -it -p 4566:4566 -p 4510-4559:4510-4559 localstack/localstack
```

Command for Docker-compose to run localstack

```docker
version: '3.6'

services:
  localstack:
    image: localstack/localstack
    container_name: localstack
    network_mode: bridge
    ports:
      - "4566:4566"
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
```

Command to create s3 bucket on localstack

```cli
aws --endpoint-url http://localhost:4566 s3 mb s3://user-uploads
```

Command to list s3 bucket on localstack

```cli
aws --endpoint-url http://localhost:4566 s3 ls
```

_______________________________________

### Some important command. 


Create sqs queue locally

```cli
aws --endpoint-url=http://localhost:4566 sqs create-queue --region=us-west-2 --queue-name trial-proj1
```

List SQS queue

```cli
aws --endpoint-url=http://localhost:4566 sqs list-queues --region=us-west-2
```

Create a SNS topic

```cli
aws --endpoint-url=http://localhost:4566 sns create-topic --region=us-west-2 --name trial-proj1-sns
```

List SNS topic

```cli
aws --endpoint-url=http://localhost:4566 sns list-topics --region=us-west-2
```

Subscribe to SNS topic

```cli
aws --endpoint-url=http://localhost:4566 sns subscribe --region=us-west-2 --topic-arn arn:aws:sns:us-west-2:000000000000:trial-proj1-sns --protocol sqs --notification-endpoint http://localhost:4566/000000000000/trial-proj1
```

List subscriptions

```cli
aws --endpoint-url=http://localhost:4566 sns list-subscriptions --region=us-west-2
```

Read the message

```cli
aws --endpoint-url=http://localhost:4566 sqs receive-message --region=us-west-2 --queue-url http://localhost:4566/000000000000/trial-proj1
```

Run go program

```cli
go run main.go -m 'MY FIRST MESSAGE' -t 'arn:aws:sns:us-west-2:000000000000:trial-proj1-sns'
```

## Useful links

*https://towardsaws.com/sns-and-sqs-with-localstack-using-golang-16b291f45e0b*

*https://learnbatta.com/blog/aws-localstack-with-docker-compose/*
