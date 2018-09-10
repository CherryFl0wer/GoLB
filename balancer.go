package main

import "fmt"

type Balancer struct {
	pool     *MinHeap
	doneChan chan *Worker
}

/*
	Initialize Pool of Worker
*/
func InitBalancer() *Balancer {

	// Create new balancer
	balancer := new(Balancer)
	balancer.pool = InitHeap(nWorker)               // init heap for the n workers
	balancer.doneChan = make(chan *Worker, nWorker) // Create workers channel

	for i := 0; i < nWorker; i++ {
		worker := new(Worker)
		worker.requests = make(chan Request, nRequester)
		worker.priority = 0

		balancer.pool.Insert(worker)

		go worker.executor(balancer.doneChan)
	}

	return balancer
}

func (b *Balancer) Balancing(requests chan Request) {
	for {
		select {
		case request := <-requests:
			b.dispatch(request)
		case worker := <-b.doneChan:
			b.completed(worker)
		}
	}
}

func (b *Balancer) dispatch(req Request) {
	work := b.pool.ExtractMin()
	work.requests <- req
	work.priority++
	b.pool.Insert(work)
}

func (b *Balancer) completed(work *Worker) {
	b.pool.SetPriority(work.index, work.priority-1)
	fmt.Println("completed request inside Worker n° : ", work.index, " priority is ", work.priority+1)
}
