package parser

import (
	"SpiderMerryU/engine"
	"regexp"
)

//<a href="/user/7386e7d179050007f16ddf25c8dc5c47.html
// <p class="info">栗子                                29岁                                <span>

var (
	profileRe = regexp.MustCompile(`<a href="(/user/[0-9a-z]+[^.html]).html`)
	namesRe   = regexp.MustCompile(` <p class="info">([\p{Han}]+|[\P{Han}]*\s)<span>`)
)

func ParserCity(contents []byte, url string) engine.ParseResult {
	//fmt.Printf("%s/n", contents)
	cityprefix := "https://www.51marryyou.com"
	matchs := profileRe.FindAllSubmatch(contents, -1)
	//namematchs := namesRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}

	//for _, namesubmatch := range namematchs {
	//	//result.Items = append(result.Items, "Person "+string(namesubmatch[1])) //用户名,用于打印输出
	//}
	for _, submatch := range matchs {
		//result.Items = append(result.Items, "Person "+string(submatch[2])) //用户名,用于打印输出
		result.Requests = append(result.Requests,
			engine.Request{
				Url:        cityprefix + string(submatch[1]) + ".html", //用户URL
				ParserFunc: ParserProfile,
			},
		)
	}

	//当前页面的相关用户列表
	citymatchs := CityList.FindAllSubmatch(contents, -1)
	for _, submatch := range citymatchs {
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
