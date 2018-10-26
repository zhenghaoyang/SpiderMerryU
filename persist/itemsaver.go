package persist

import (
	"SpiderMerryU/engine"
	"context"
	"errors"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"gopkg.in/olivere/elastic.v5/config"
	"log"
)

func ItemSaver(index string) chan engine.Item {
	out := make(chan engine.Item)
	//SetSniff 客户端维护集群转态,跑在docker内网，没办法sniff
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
		fmt.Printf("err = %s\n", err)
		panic(err)
	}

	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Got Item #%d: %v", itemCount, item)
			itemCount++
			//log.Printf("当前保存信息为%+v", item)
			err := save(client, index, item)
			if err != nil {
				log.Println("save item 存错信息为 = ", err)
				continue
			}
		}
	}()
	return out
}

func save(client *elastic.Client, index string, item engine.Item) error {

	if item.Type == "" {
		return errors.New("Must supply Type")
	}
	indexService := client.Index().Index(index).
		Type(item.Type).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}

	//_, err := client.Index().Index(index).Type("user").
	//	BodyJson(item).Do(context.Background())
	//fmt.Printf("resp %+v\n", resp)
	return err
}
