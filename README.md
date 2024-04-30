# Run GO app
go run main.go

## Make sure you have GO cli installed
[You can get it here](https://go.dev/doc/install)

## Lambda functions 
Lambda functions are defined in the lambda package.
Each should have its own main package in order to build the executable.

To build a Lambda function, please run the following command
`go run ./lambdas/{packagename}/main.go`

Note the leading `./` on the command line, it is important!