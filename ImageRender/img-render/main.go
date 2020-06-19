package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type object struct {
}

type shelf struct {
	obj []object
}

// configer is JSONからレンダリングに関する指定を取得し、各値を設定する
func configer(path string, s *shelf) error {
	return errors.New("aaaa")
}

func handler(ctx context.Context, s3Event events.S3Event) {
	for _, record := range s3Event.Records {
		if record.EventName == "ObjectCreated:Put" {
			s3 := record.S3
			fmt.Printf("[SUCCESS] event doing!! - %s \n", record.EventName)
			fmt.Printf("[%s - %s] Bucket = %s, Key = %s \n", record.EventSource, record.EventTime, s3.Bucket.Name, s3.Object.Key)
		} else {
			fmt.Printf("[Skip] event doing!! - %s \n", record.EventName)
		}
	}
}

func main() {
	lambda.Start(handler)
}
