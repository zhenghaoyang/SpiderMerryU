package engine

import (
	"SpiderMerryU/fetcher"
	"fmt"
	"log"
	"os"
)

func Run(seeds ...Request) {
	var requests []Request
	//收集种子
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		//取出种子
		r := requests[0]
		//更新
		requests = requests[1:]
		//取得URL的文本数据
		log.Printf("now is Fetching URL :%s", r.Url)
		content, err := fetcher.Fetch(r.Url)
		if err != nil {
			continue
		}
		//调用解析器，取得匹配的数据
		//fmt.Printf("Start ParserFunc----------------------\n")
		//fmt.Printf("Start ParserFunc %v\n", r.Url)
		//fmt.Printf("Start ParserFunc----------------------\n")

		if r.ParserFunc == nil {
			fmt.Printf("Done %+v", r.ParserFunc)
			os.Exit(1)
		}

		parseResult := r.ParserFunc(content,"")
		//下一级的request添加到种子队列
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("got item %v\n", item)
		}
		//打印城市URL
		for _, req := range parseResult.Requests {
			log.Printf("got item  %s\n", req.Url)
		}

	}

}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}
