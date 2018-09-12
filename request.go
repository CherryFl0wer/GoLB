package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/golang-collections/collections/queue"
)

type FunctionCallType func()
type Request struct {
	operation FunctionCallType // An action to do
	response  chan bool        // The response to the action
}

var requestQueueManager *queue.Queue = queue.New()

var httpClient *http.ServeMux = http.NewServeMux()

// AddRequestToQueue  adding request to queue
// ----
func AddRequestToQueue(operation FunctionCallType) {
	requestQueueManager.Enqueue(operation)
}

// Requester  Request specific method
// ----
func Requester(operation http.HandlerFunc, requests chan<- Request) {
	resultChan := make(chan bool)
	for {
		time.Sleep(time.Duration(rand.Int63n(int64(time.Millisecond)))) // Simulate load, wait before next request
		req := requestQueueManager.Dequeue()
		if req == nil {
			continue
		}
		r := req.(FunctionCallType)
		requests <- Request{r, resultChan}

		<-resultChan
	}

}
