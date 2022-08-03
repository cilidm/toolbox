package pool

import "sync"

type Pool struct {
	queue chan int
	wg    *sync.WaitGroup
}

func NewPool(size int) *Pool {
	if size <= 0 {
		size = 1
	}
	return &Pool{queue: make(chan int, size), wg: &sync.WaitGroup{}}
}

func (p *Pool) Add(size int) {
	for i := 0; i < size; i++ { // size > 0
		p.queue <- 1
	}
	for i := 0; i > size; i-- { // size < 0
		<-p.queue
	}
	p.wg.Add(size)
}

func (p *Pool) Done() {
	<-p.queue
	p.wg.Done()
}

func (p *Pool) Wait() {
	p.wg.Wait()
}
