package serverless

const ServerlessGoLambda string = `package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/icarus-sullivan/mock-lambda"
)

func hello(context.Context, events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: "Hello from λ!"}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(hello)
}
`
const ServerlessGoTest string = `package {{ .FunctionName }}_test

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	type Request struct {
		name string
	}
	tests := []struct {
		request events.APIGatewayProxyRequest
		expect  string
		err     error
	}{
		{
			request: events.APIGatewayProxyRequest{name: "Developer"},
			expect:  "Hello Developer from λ!",
			err:     nil,
		},
		{
			request: events.APIGatewayProxyRequest{name: ""},
			expect:  "Hello from λ!",
			err:     {{ .FunctionName }}.ErrDataNotProvided,
		},
	}

	for _, test := range tests {
		response, err := {{ .FunctionName }}.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, response.Body)
	}

}
`
