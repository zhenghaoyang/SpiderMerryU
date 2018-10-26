package parser

import (
	"SpiderMerryU/engine"
	"regexp"
)

//([^.html]+)
//([^>]+)
//<span><a href="/seekMarry/110000_0_8_1.html" target='_blank'>北京</a></span>
var (
	CityList  = regexp.MustCompile(`<span><a href="(/seekMarry/[^.html]+).html" target='_blank'>([^>]+)</a></span>`)

	//下一页
//<a href="/seekMarry/350200_0_8_3.html">
)

func ParserCityList(contents []byte,url string) engine.ParseResult {
	matchs := CityList.FindAllSubmatch(contents, -1)
	cityprefix := "https://www.51marryyou.com"

	result := engine.ParseResult{}
	for _, submatch := range matchs {
		//result.Items = append(result.Items, string(submatch[2])) //城市名,用于打印输出
		//下一级Request
		result.Requests = append(result.Requests,
			engine.Request{
				Url:        cityprefix + string(submatch[1]), //城市URL
				ParserFunc: ParserCity,
			},
		)
	}








	return result
}
