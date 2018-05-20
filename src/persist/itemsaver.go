package persist

import (
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"context"
)

func Itemsaver()chan interface{} {
	out:=make(chan interface{})
	go func() {
		itemcount:=0
		for  {
			item:=<-out

			itemcount++
			fmt.Printf("id is itemsaver got the item---%d:%+v \n",itemcount,item)
			save(item)
			//if err!=nil{
			//	continue
			//}
			//fmt.Printf("id is %s,itemsaver got the item---%d:%v \n",id,itemcount,item)

		}
	}()
	return out
}
func save(item interface{}) (id string,err error) {
	client,err:=elastic.NewClient(
		elastic.SetSniff(false))
	if err!=nil{
		fmt.Printf("%d",11)
		panic(err)
	}
	fmt.Printf("333")
	fmt.Printf("id is itemsaver got the item---:%+v \n",item)
	fmt.Printf("444")
	resp,err:=client.Index().Index("dataing_profile").Type("zhenai").BodyJson(item).Do(context.Background())
	if err!=nil{
		fmt.Printf("%d",222)
		panic(err)
		return "",err
	}
	fmt.Printf("%+v",resp)
	return resp.Id,nil
}
