package engine

import (
	"fetch"
	"log"
)

type Simpleengine struct {

}
func (e Simpleengine) Run(seed ...Request)  {
	var requests []Request
	for _,r:=range seed{
		requests=append(requests,r)
	}
	for len(requests)>0 {
		r:=requests[0]
		requests=requests[1:]
		parseresult,err:=Worker(r)
		if err!=nil{
			continue
		}
		requests=append(requests,parseresult.Requestes...)
		for _,item:=range parseresult.Items {
			log.Printf("got item %s\n",item)
		}
	}
}

func Worker(r Request)(ParseResult,error) {
	body,err:=fetch.Fetch(r.Url)
	if err!=nil{

		return ParseResult{},err
	}
	parseresult:=r.ParserFunc(body)
	return parseresult,nil

}