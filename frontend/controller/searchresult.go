package controller

import (
	"SpiderMerryU/constvar"
	"SpiderMerryU/engine"
	"SpiderMerryU/frontend/model"
	"SpiderMerryU/frontend/view"
	"context"
	"gopkg.in/olivere/elastic.v5"
	"gopkg.in/olivere/elastic.v5/config"
	"net/http"
	"reflect"
	"regexp"

	"strconv"
	"strings"
)

type SearchResultHandler struct {
	view   view.SerachResultView
	client *elastic.Client
}

func CreateSearchResultHandler(filename string) SearchResultHandler {
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
		panic(err)
	}

	return SearchResultHandler{
		view:   view.CreateSearchResultView(filename),
		client: client,
	}
}

//type Handler interface {
//	ServeHTTP(ResponseWriter, *Request)
//}
//ServeHTTP 是实现 handle的接口方法 WTF
func (handler SearchResultHandler) ServeHTTP(
	w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}

	var page model.SearchResult
	page, err = handler.getSearchResult(q, from)
	err = handler.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	//fmt.Fprintf(w, "q = %s from = %d", q, from)
}
func (handler SearchResultHandler) getSearchResult(
	q string, from int) (model.SearchResult, error) {

	var result model.SearchResult
	result.Query = q

	resp, err := handler.client.Search(constvar.Db).
		Query(elastic.NewQueryStringQuery(rewriterquerystring(q))).
		From(from).
		Do(context.Background())

	if err != nil {
		return result, err
	}
	result.Hits = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)

	return result, nil
}

func rewriterquerystring(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	return re.ReplaceAllString(q, "Payload.$1:")
}
