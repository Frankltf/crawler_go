package Parser

import (
	"testing"
	"io/ioutil"
	"fmt"
)

func TestPareserPersonlist(t *testing.T)  {
	contents,err:=ioutil.ReadFile("profilelist.html")
	if err!=nil{
		panic(err)
	}
	result:=PrintPersonlist(string(contents))
	fmt.Printf("%v\n",result)
	//expecturls:=[]string{"http://city.zhenai.com/aba","http://city.zhenai.com/akesu","http://city.zhenai.com/alashanmeng",}
	//expectcitys:=[]string{"阿坝","阿克苏","阿拉善盟",}
	//for i,url:=range expecturls {
	//	if result.Requestes[i].Url!=url{
	//		t.Errorf("expected url %d is %s ,but is %s",i,url,result.Requestes[i].Url)
	//	}
	//}
	//for i,city:=range expectcitys {
	//	if result.Items[i].(string)!=city{
	//		t.Errorf("expected city %d is %s ,but is %s",i,city,result.Items[i])
	//	}
	//}
	//const resultsize  = 470
	//if len(result.Requestes)!=resultsize{
	//	t.Errorf("should have %d request ,but have %d",resultsize,len(result.Requestes))
	//}
	//if len(result.Items)!=resultsize{
	//	t.Errorf("should have %d request ,but have %d",resultsize,len(result.Items))
	//}
}
