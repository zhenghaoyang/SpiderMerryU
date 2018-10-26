package view

import (
	"SpiderMerryU/frontend/model"
	"html/template"
	"io"
)

type SerachResultView struct {
	template *template.Template
}

func CreateSearchResultView(filename string) SerachResultView {
	return SerachResultView{
		template: template.Must(template.ParseFiles(filename)),
	}
}

func (s SerachResultView) Render(w io.Writer, data model.SearchResult) error {
	return s.template.Execute(w, data)
}
