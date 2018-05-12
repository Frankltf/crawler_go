package main

import (
	"engine"
	"zhenai/Parser"
)

func main() {
	engine.Run(engine.Request{
		"http://city.zhenai.com/",
		Parser.PrintCityList,
	})
}