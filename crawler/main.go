package main

import (
	"engine"
	"zhenai/Parser"
	"scheduler"
)

func main() {
	e:=engine.Concurrentengine{
		Scheduler:&scheduler.QueuedScheduler{},
		WorkerCount:1000,
	}
	e.Run(engine.Request{
		"http://city.zhenai.com/",
		Parser.PrintCityList,
	})
}