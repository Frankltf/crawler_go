package Parser

import (
	"regexp"
	"engine"
)

const personlistregexp  = `<td><a href="(http://album.zhenai.com/u/[\d]+)" target="_blank">.*</a></td>`
func PrintPersonlist(content string) engine.ParseResult {
	re:=regexp.MustCompile(personlistregexp)
	matches:=re.FindAllStringSubmatch(content,-1)
	result:=engine.ParseResult{}
	for _,m :=range matches  {
		result.Requestes=append(result.Requestes,engine.Request{
			m[1],
			func(str string) engine.ParseResult {
				return ParseProfile(str)
			},
		})
	}
	return result
}