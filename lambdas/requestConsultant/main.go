package main

import (
	"errors"
	"goAngularTryout/lambdas/utils"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"fmt"
	"log"
)

type Event struct {
	Name string `json:"Name"`
}

type Consultant struct {
	Name string
}

type MyResponse struct {
	Tables string `json:"Tables"`
}

func HandleLambdaEvent(event *Event) (*MyResponse, error) {
	dbClient := utils.CreateDbSession()

	tableName := "Consultants"
	result, err := dbClient.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Name": {
				S: aws.String(event.Name),
			},
		},
	})

	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
	}

	if result.Item == nil {
		msg := "Could not find consultant with name '" + event.Name + "'"
		return nil, errors.New(msg)
	}

	consultant := Consultant{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &consultant) // Map resultitem to consultant variable
	if err != nil {
		panic(err) // PANIC! This is a programming error so we stop execution
	}

	return &MyResponse{Tables: fmt.Sprint(consultant)}, nil

}

func main() {
	lambda.Start(HandleLambdaEvent)
}
