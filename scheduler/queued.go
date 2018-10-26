package scheduler

import (
	"SpiderMerryU/engine"
)

type QueuedScheduler struct {
	RequestChan chan engine.Request
	WorkerChan  chan chan engine.Request ////每个worker 对应的不同的chan,队列实现
}

func (s *QueuedScheduler) WorkChan() chan engine.Request {
	//每个worker对应自己的一个chan
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.RequestChan <- r
}

//确定有一个work准备好了，可以接收request

func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.WorkerChan <- w
}

func (s *QueuedScheduler) Run() {
	s.WorkerChan = make(chan chan engine.Request)
	s.RequestChan = make(chan engine.Request)

	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			//当对列有数据,发送数据
			if len(requestQ) > 0 && len(workerQ) > 0 {
				//不能再这里发送对列数据,会阻塞，不能更新队列数据，因为后面的select不一定选中
				activeRequest = requestQ[0]
				activeWorker = workerQ[0] //activeWorker本身是个chan
			}
			select {
			case r := <-s.RequestChan:
				//	send r to 那个worker (未知) 解决加进队列
				requestQ = append(requestQ, r) //排队
			case w := <-s.WorkerChan:
				//	send request_next(未知) to worker
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest: //确实选中，才能更新队列数据
				// 当为activeWorkernil不会被选中
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}

		}

	}()
}
