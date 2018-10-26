package engine

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan Item
}

type Scheduler interface {
	Submit(Request)
	WorkChan() chan Request
	ReadyNotifier
	Run()
}
type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	//配置worker的输入
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		//创建worker 需要两个chan,传递数据text in content out
		createWorker(e.Scheduler.WorkChan(), out, e.Scheduler)
	}
	for _, r := range seeds {
		//requests = append(requests, r)
		//分发Request
		e.Scheduler.Submit(r)
	}

	//收worker的数据,解析后的数据
	for {
		result := <-out //循环等待
		for _, item := range result.Items {
			go func() { //传送用户信息
				e.ItemChan <- item
			}()

		}
		//把request 送给request.Requests,往下分发
		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}

	}

}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {

	go func() {
		for {
			//告诉schedule work准备好了，
			ready.WorkerReady(in)
			//文本数据
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			//解析完的数据
			out <- result
		}
	}()
}
