package main

import (
	"math/rand"
	"time"
)

type OperationResponseType int
type OperationMethod func() OperationResponseType

type Request struct {
	operation OperationMethod            // An action to do
	response  chan OperationResponseType // The response to the action
}

func Requester(operation OperationMethod, requests chan<- Request) {
	resultChan := make(chan OperationResponseType)
	for {
		time.Sleep(time.Duration(rand.Int63n(int64(time.Millisecond)))) // Simulate load, wait before next request
		requests <- Request{operation, resultChan}
		sendResult(<-resultChan)
	}

}

func sendResult(r OperationResponseType) {
	//fmt.Println("Result ", r)
}
