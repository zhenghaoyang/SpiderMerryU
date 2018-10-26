package engine

import (
	"SpiderMerryU/fetcher"
	"log"
)

//处理抓取数据与解析数据，返回解析的结果
func worker(r Request) (ParseResult, error) {
	//取得URL的文本数据
	log.Printf("now is Fetching URL :%s", r.Url)
	//开始真正的抓取数据
	body, err := fetcher.Fetch(r.Url)

	if err != nil {
		log.Printf("Fether:error:"+
			"fetching url %s %v", r.Url, err)
		return ParseResult{}, err
	}

	//调用解析器，取得匹配的数据

	//fmt.Printf("Start ParserFunc %v\n", r.Url)

	return r.ParserFunc(body,r.Url), nil
}