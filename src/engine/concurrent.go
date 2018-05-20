package engine

import (

)

type Concurrentengine struct {
	Scheduler Scheduler
	WorkerCount int
	Itemchan chan interface{}
}
type Scheduler interface {
	Submit(Request)
	ConfigureWorkchan(chan Request)
	WorkReady(chan Request)
	Run()
} 
func (e *Concurrentengine) Run(seeds ...Request)  {
	out:=make(chan ParseResult)
	e.Scheduler.Run()
	for i:=0;i<e.WorkerCount;i++ {
		createworker(out,e.Scheduler)
	}
	for _,r := range seeds{
		e.Scheduler.Submit(r)
	}
	for {
		result:=<-out
		for _,item:=range result.Items{
			go func() {e.Itemchan<-item}()
		}
		for _,request :=range result.Requestes{
			e.Scheduler.Submit(request)
		}
	}
}
func createworker(out chan ParseResult,s Scheduler)  {
	in:=make(chan Request)
	go func() {
		for  {
			s.WorkReady(in)
			request:=<-in
			parseresult,err:=Worker(request)
			if err!=nil {
				continue
			}
			out<-parseresult
		}
	}()
}