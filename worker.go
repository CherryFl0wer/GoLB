package main

// Structure used for the heap
type Worker struct {
	requests chan Request
	index    int
	priority int
}

func (w *Worker) executor(done chan *Worker) {
	for {
		request := <-w.requests
		request.operation()
		request.response <- true
		done <- w
	}
}
