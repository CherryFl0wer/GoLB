package main

import (
	"fmt"
	"math/rand"
	"time"
)

type OperationResponseType int
type OperationMethod func() OperationResponseType // Can add argument such as http request

type Request struct {
	operation OperationMethod            // An action to do
	response  chan OperationResponseType // The response to the action
}

func requester(operation OperationMethod, requests chan<- Request) {
	resultChan := make(chan OperationResponseType)
	for {
		time.Sleep(time.Duration(rand.Int63n(int64(time.Millisecond)))) // Simulate load, wait before next request
		requests <- Request{operation, resultChan}
		sendResult(<-resultChan)
	}
}

func sendResult(r OperationResponseType) {
	fmt.Println("Result %d", r)
}
