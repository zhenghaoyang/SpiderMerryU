package persist

import (
	"SpiderMerryU/model"
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"gopkg.in/olivere/elastic.v5/config"
	"testing"
)

func TestItemSaver(t *testing.T) {
	expected := model.Profile{
		Name:       "惠儿",
		Age:        50,
		Height:     156,
		Weight:     0,
		Income:     "3000元以下",
		Gender:     "女",
		Xinzuo:     "魔羯座",
		Marriage:   "离异",
		Education:  "高中及以下",
		Occupation: "销售总监",
		House:      "租房",
		Car:        "未购车",
	}

	var flag = false
	cfg := &config.Config{
		URL:         "http://114.116.53.46:9200",
		Sniff:       &flag,
		Healthcheck: &flag,
		Shards:      1,
		Replicas:    0,
	}
	client, err := elastic.NewClientFromConfig(cfg)
	if err != nil {
		fmt.Printf("err = %+v\n", err)
		panic(err)
	}

	err = save(client, "merryu", expected)

	resp, err := client.Get().
		Index("merryu").
		Type("user").Id("AWaU-Cj2D_ljr9OANrlE").
		Do(context.Background())
	var actual model.Profile
	err = json.Unmarshal(*resp.Source, &actual)
	fmt.Printf("user %+v", actual)

}

type Config struct {
	URL         string
	Index       string
	Username    string
	Password    string
	Shards      int
	Replicas    int
	Sniff       *bool
	Healthcheck *bool
	Infolog     string
	Errorlog    string
	Tracelog    string
}
