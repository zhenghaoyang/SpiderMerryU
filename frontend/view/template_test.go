package view

import (
	"SpiderMerryU/engine"
	"SpiderMerryU/frontend/model"
	common "SpiderMerryU/model"
	"os"
	"testing"
)

func TestSerachResultView_Render(t *testing.T) {

	//template := template.Must(template.ParseFiles("template.html"))
	//template.Execute(os.Stdout, page)
	//template.Execute(out, page)
	view := CreateSearchResultView("template.html")

	out, _ := os.Create("template_test.html")

	page := model.SearchResult{}

	page.Hits = 100
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/107906650",
		Type: "zhenai",
		Id:   "107906650",
		Payload: common.Profile{
			Name:       "惠儿",
			Age:        25,
			Height:     156,
			Weight:     0,
			Income:     "3000元以下",
			Gender:     "女",
			Marriage:   "离异",
			Education:  "高中及以下",
			Occupation: "销售总监",
			Hokou:      "四川阿坝",
			Xinzuo:     "魔羯座",
			House:      "租房",
			Car:        "未购车",
		},
	}

	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	//执行数据匹配
	view.Render(out, page)

}
