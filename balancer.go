package main

type Balancer struct {
	pool     *MinHeap
	doneChan chan *Worker
}

/*
	Initialize Pool of Worker
*/
func InitBalancer(nbWorker, nbRequests int) *Balancer {

	// Create new balancer
	balancer := new(Balancer)
	balancer.pool = InitHeap(nbWorker)               // init heap for the n workers
	balancer.doneChan = make(chan *Worker, nbWorker) // Create workers channel

	for i := 0; i < nbWorker; i++ {
		worker := new(Worker)
		worker.requests = make(chan Request, nbRequests)
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
}
