package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(HandleRequest)
}

type Event struct {
	Id string
}

func HandleRequest(ctx context.Context, event Event) (string, error) {
	fmt.Println(event.Id)
	return event.Id, nil
}
