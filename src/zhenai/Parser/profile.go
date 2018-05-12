package Parser

import (
	"regexp"
	"model"
	"engine"
)

const profileAge  =`<td><span class="label">年龄：</span>([\d]+)岁</td>`
const profileHeight  =`<td><span class="label">身高：</span>([\d]+)CM</td>`
const profileEducation  =`<td><span class="label">学历：</span>([^>]+)</td>`
const profileName  =`<h1 class="ceiling-name ib fl fs24 lh32 blue">([^>]+)</h1>`
func ParseProfile(content string)engine.ParseResult  {
	profile:=model.Profile{}
	profile.Age=finddata(profileAge,content)

	profile.Height=finddata(profileHeight,content)

	profile.Education=finddata(profileEducation,content)

	profile.Name=finddata(profileName,content)


	result:=engine.ParseResult{
		Items:[]interface{}{profile},
	}
	return result
}
func finddata(reg string,content string)string  {
	re:=regexp.MustCompile(reg)
	matches:=re.FindStringSubmatch(content)
	var datacell string
	if matches!=nil{
		datacell=matches[1]
	}
	return datacell
}
