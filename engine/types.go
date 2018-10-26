package engine

type Request struct {
	Url        string
	ParserFunc ParserFun
}

type ParserFun func(
	contents []byte, url string) ParseResult

type ParseResult struct {
	Requests []Request     //城市URL
	Items    []Item //城市名
}

type Item struct {
	Url     string
	Id      string
	Type    string
	Payload interface{}
}
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
