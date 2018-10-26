package parser

import (
	"io/ioutil"
	"testing"
)

func TestCityList(t *testing.T) {

	contents, _ := ioutil.ReadFile("citylist.html")
	result := ParserCityList(contents)
	const resultSize = 35
	if len(result.Requests) != resultSize {
		t.Errorf("Result should have %d "+""+
			"requests; but had %d", resultSize, len(result.Requests))
	}
	if len(result.Items) != resultSize {
		t.Errorf("Result should have %d "+""+
			"requests; but had %d", resultSize, len(result.Items))
	}

	expectedUrls := []string{
		"https://www.51marryyou.com/seekMarry/110000_0_8_1", "https://www.51marryyou.com/seekMarry/310000_0_8_1", "https://www.51marryyou.com/seekMarry/120000_0_8_1",
	}
	expectedCities := []string{
		"北京", "上海", "天津",
	}
	//测试少量数据
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but "+
				"was %s", i, url, result.Requests[i].Url)
		}
	}
	for i, city := range expectedCities {
		//明确知道string 使用类型断言
		if result.Items[i].(string) != city {
			t.Errorf("expected city #%d: %s; but "+
				"was %s", i, city, result.Items[i])
		}
	}
}
