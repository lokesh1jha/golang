package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"what is your name?"`
	Age  int    `json:"how old are you?"`
}

type MyResponse struct {
	Message string `json:"Answer:"`
}

func HandleRequest(event MyEvent) (MyResponse, error) {
	return MyResponse{Message: fmt.Sprintf("Hello %s, you are %d years old", event.Name, event.Age)}, nil // nil is for error
}

func main() {
	fmt.Println("Starting Lambda")
	lambda.Start(HandleRequest)
}
