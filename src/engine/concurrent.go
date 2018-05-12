package engine

import (
	"fmt"
)

type Concurrentengine struct {
	Scheduler Scheduler
	WorkerCount int
}
type Scheduler interface {
	Submit(Request)
	ConfigureWorkchan(chan Request)
} 
func (e *Concurrentengine) Run(seeds ...Request)  {
	in:=make(chan Request)
	out:=make(chan ParseResult)
	e.Scheduler.ConfigureWorkchan(in)
	for i:=0;i<e.WorkerCount;i++ {
		createworker(in,out)
	}
	for _,r := range seeds{
		e.Scheduler.Submit(r)
	}
	itemcount:=0
	for {
		result:=<-out
		for _,item:=range result.Items{
			fmt.Printf("got item (itemcount is :%d): %v \n",itemcount,item)
			itemcount++
		}
		for _,request :=range result.Requestes{
			e.Scheduler.Submit(request)
		}
	}
}
func createworker(in chan Request,out chan ParseResult)  {
	go func() {
		for  {
			request:=<-in
			parseresult,err:=Worker(request)
			if err!=nil {
				continue
			}
			out<-parseresult
		}
	}()
}