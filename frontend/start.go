package main

import (
	"SpiderMerryU/frontend/controller"
	"html/template"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(
		http.Dir("../SpiderMerryU/frontend/view/"),
	))

	http.Handle("/_search", controller.
		CreateSearchResultHandler("../SpiderMerryU/frontend/view/template.html"))
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}

}
func indexHandler(w http.ResponseWriter, req *http.Request) {
	template.Must(template.ParseFiles("c:/goproject/src/SpiderMerryU/frontend/view/index.html")).Execute(w, req)

}
