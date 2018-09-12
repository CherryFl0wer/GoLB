package main

import (
	"fmt"
	"runtime"
)

const nWorker = 10
const nRequester = 100

var MaxNbThread int = runtime.NumCPU()

func main() {
	runtime.GOMAXPROCS(MaxNbThread)

	requests := make(chan Request)
	for i := 0; i < nRequester; i++ {
		AddRequestToQueue(func() {
			fmt.Println("You're lost, go home")
		})

		go Requester(requests)
	}

	InitBalancer().Balancing(requests)

}
