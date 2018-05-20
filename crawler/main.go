package main

import (
	"engine"
	"zhenai/Parser"
	"scheduler"
	"persist"
)

func main() {
	e:=engine.Concurrentengine{
		Scheduler:&scheduler.QueuedScheduler{},
		WorkerCount:1000,
		Itemchan:persist.Itemsaver(),
	}
	e.Run(engine.Request{
		"http://city.zhenai.com/",
		Parser.PrintCityList,
	})
}