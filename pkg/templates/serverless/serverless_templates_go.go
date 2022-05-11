package serverless

const ServerlessGoLambda string = `package main

import (
        "fmt"
        "context"
        "github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
        Name string ` + `json:"name"` + `
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
        return fmt.Sprintf("Hello %s!", name.Name ), nil
}

func main() {
        lambda.Start(HandleRequest)
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
			expect:  "Hello Developer",
			err:     nil,
		},
		{
			request: events.APIGatewayProxyRequest{name: ""},
			expect:  "",
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
