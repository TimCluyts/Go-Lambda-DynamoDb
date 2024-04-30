package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"Name"`
}

type MyResponse struct {
	Message string `json:"Result"`
}

/**
* NOTE: GO can return multiple responses from a single function
*		Here for example we return a response AND an error
**/
func HandleLambdaEvent(event *MyEvent) (*MyResponse, error) {
	if event == nil {
		return nil, fmt.Errorf("received nil event")
	}
	return &MyResponse{Message: fmt.Sprintf("Hello there, %s", event.Name)}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
