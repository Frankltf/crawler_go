package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"github.com/axgle/mahonia"
	"regexp"
)

func main() {
	resp,err:=http.Get("http://city.zhenai.com/")
	if(err!=nil){
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode==http.StatusOK{
		all,err:=ioutil.ReadAll(resp.Body)
		if err!=nil{
			panic(err)
		}
		newall:=ConverToString(string(all),"gbk","utf-8")
		//fmt.Printf("%s\n",newall)
		PrintCityList(newall)
	}else {
		fmt.Println("Error:status code",resp.StatusCode)
		return
	}
}
func PrintCityList(content string)  {
	re:=regexp.MustCompile(`<a href="(http://city.zhenai.com/[a-zA-Z0-9]+)"[^>]*>(.*)</a>`)
	matches:=re.FindAllStringSubmatch(content,-1)
	for _,m :=range matches  {
		fmt.Printf("city:%s ,url:%s \n",m[2],m[1])
	}
}
func ConverToString(src string,srccode string,tagcode string)string  {
	srccoder:=mahonia.NewDecoder(srccode)
	srcResult:=srccoder.ConvertString(src)
	tagcoder:=mahonia.NewDecoder(tagcode)
	_,data,_:=tagcoder.Translate([]byte(srcResult),true)
	result:=string(data)
	return result
}
