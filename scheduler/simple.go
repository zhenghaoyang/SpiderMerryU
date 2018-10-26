package scheduler

import "SpiderMerryU/engine"

type SimpleScheduler struct {
	WorkerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureWorkerChan(c chan engine.Request) {
	s.WorkerChan = c
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	go func() {
		s.WorkerChan <- request
	}()
}
