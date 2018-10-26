package main

import (
	"SpiderMerryU/constvar"
	"SpiderMerryU/engine"
	"SpiderMerryU/parser"
	"SpiderMerryU/persist"
	"SpiderMerryU/scheduler"
)

func main() {
	//contents, _ := fetcher.Fetch("https://www.51marryyou.com/search.html")

	e := &engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: constvar.WorkerCount,
		ItemChan:    persist.ItemSaver(constvar.Db),
	}
	e.Run(engine.Request{
		Url:        constvar.Index,
		ParserFunc: parser.ParserCityList,
	})

	//e.Run(engine.Request{
	//	Url:        "https://www.51marryyou.com/seekMarry/350100_0_8_1.html",
	//	ParserFunc: parser.ParserCity,
	//})

	//个人详情页
	//https://www.51marryyou.com/user/de7c99cba7b103b8e7e7252fab5d51db.html
	const person = "https://www.51marryyou.com/user/de7c99cba7b103b8e7e7252fab5d51db.html"
	e.Run(engine.Request{
		Url:        person,
		ParserFunc: parser.ParserProfile,
	})

}
