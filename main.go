package main

import (
	"runtime"
)

const nWorker = 10
const nRequester = 100

const MaxNbThread = 4 //runtime.NumCPU()

func main() {
	runtime.GOMAXPROCS(MaxNbThread)

	requests := make(chan Request)
	for i := 0; i < nRequester; i++ {
		go Requester(func() OperationResponseType {
			return 1
		}, requests)
	}

	InitBalancer().Balancing(requests)
}
