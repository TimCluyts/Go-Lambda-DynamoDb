package main

import (
	"goAngularTryout/lambdas/utils"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"fmt"
)

type MyResponse struct {
	Tables string `json:"Tables"`
}

/**
* NOTE: GO can return multiple responses from a single function
*		Here for example we return a response AND an error
**/
func HandleLambdaEvent() (*MyResponse, error) {
	dbClient := utils.CreateDbSession()

	// Input configuration - here we will use blank to get all tables back
	input := &dynamodb.ListTablesInput{}

	var allTableNames []string

	for { // Infinite loop until we break
		result, err := dbClient.ListTables(input)

		if err != nil {
			return nil, err
		}

		// First variable here is ALWAYS index, which we do not need
		// So for GO we ignore variables by using underscore
		for _, n := range result.TableNames {
			allTableNames = append(allTableNames, *n)
		}

		// assign the last read tablename as the start for our next call to the ListTables function
		// the maximum number of table names returned in a call is 100 (default), which requires us to make
		// multiple calls to the ListTables function to retrieve all table names
		input.ExclusiveStartTableName = result.LastEvaluatedTableName

		if result.LastEvaluatedTableName == nil { // We reach the end of the list of tables
			break
		}
	}

	return &MyResponse{Tables: fmt.Sprint(allTableNames)}, nil

}

func main() {
	lambda.Start(HandleLambdaEvent)
}
