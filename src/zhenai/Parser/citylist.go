package Parser

import (
	"regexp"
	"engine"
)

const cityregexp  = `<a href="(http://city.zhenai.com/[a-zA-Z0-9]+)"[^>]*>(.*)</a>`
func PrintCityList(content string) engine.ParseResult {
	re:=regexp.MustCompile(cityregexp)
	matches:=re.FindAllStringSubmatch(content,-1)
	result:=engine.ParseResult{}
	limit:=10
	for _,m :=range matches  {
		result.Items=append(result.Items,m[2])
		result.Requestes=append(result.Requestes,engine.Request{
			m[1],
			func(str string) engine.ParseResult {
				return PrintPersonlist(str)
			},
		})
		limit--
		if limit<=0{
			break
		}
	}
	return result
}