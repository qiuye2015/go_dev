package scheduler

import "github.com/qiuye2015/go_dev/crawler/engine"

type SimpleScheduler struct {
	workChan chan engine.Request
}

func (s *SimpleScheduler) Run() {
	s.workChan = make(chan engine.Request)
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() { //queued中无goroutine
		s.workChan <- r
	}()
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workChan
}

func (s *SimpleScheduler) WorkerReady(requests chan engine.Request) {
}
