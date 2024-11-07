// main handles Lambda logic.
package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bartlomiej-jedrol/pr11-bowling-exercise/bowling"
)

var ErrorInternalServerError error = errors.New("internal server error")

// buildResponseBody builds API Gateway response body.
func buildResponseBody(body any) string {
	switch v := body.(type) {
	case error:
		return fmt.Sprintf(`{"error":"%v"}`, v)
	case string:
		return v
	default:
		return ""
	}
}

// buildAPIResponse builds API Gateway response.
func buildAPIResponse(statusCode int, body any) (*events.APIGatewayProxyResponse, error) {
	log.Printf("INFO: buildAPIResponse - building API Gateway response")

	resp := &events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		StatusCode: statusCode,
	}
	resp.Body = buildResponseBody(body)
	return resp, nil
}

// HandleRequest handles Bowling Game logic.
func HandleRequest(
	request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	body, err := bowling.PlayGame()
	if err != nil {
		return buildAPIResponse(http.StatusInternalServerError, ErrorInternalServerError)
	}
	return buildAPIResponse(http.StatusOK, body)
}

func main() {
	lambda.Start(HandleRequest)
}
