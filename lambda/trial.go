package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	fmt.Println("executing")
	jsonBytes, err := ioutil.ReadFile("mappings.json")
	fmt.Println(string(jsonBytes))
	fmt.Println(err)
	return fmt.Sprintf("Hello %s!", name.Name), nil
}

func main() {
	lambda.Start(HandleRequest)
}
