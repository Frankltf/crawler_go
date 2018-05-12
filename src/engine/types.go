package engine

type Request struct {
	url string
	ParserFunc func(str string) ParseResult
}
type ParseResult struct {
	Requestes []Request
	Items []interface{}
}

func NilParse(str string)ParseResult  {
	return ParseResult{}
}
