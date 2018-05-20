package fetch

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"github.com/axgle/mahonia"
)

func Fetch(url string)(string,error)  {
	resp,err:=http.Get(url)
	if(err!=nil){
		return "",err
	}
	defer resp.Body.Close()
	if resp.StatusCode==http.StatusOK{
		all,err:=ioutil.ReadAll(resp.Body)
		if err!=nil{
			panic(err)
		}
		newall:=ConverToString(string(all),"gbk","utf-8")
		return newall,nil

	}else {
		return "",fmt.Errorf("Error:status code %d:",resp.StatusCode)
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