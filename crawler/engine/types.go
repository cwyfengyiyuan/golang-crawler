package engine

type ParserFunc func(contents []byte, url string) ParseResult
type Request struct {
	Url string
	ParserFunc ParserFunc
}
type ParseResult struct {
	Requests []Request
	Items []Item
	//Items []interface{}
}
type Item struct {
	Id string
	Type string
	Url string
	Payload interface{}
}
func NilParser([]byte) ParseResult{
	return ParseResult{}
}