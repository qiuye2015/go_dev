package main

import (
	"crawler/douban"
	"crawler/engine"
	"crawler/persist"
	"crawler/scheduler"
	"crawler_distributed/config"
)

func main() {
	//engine.SimpleEngine.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: citylist.ParseCityList,
	//})

	//e1 := engine.ConcurrentEngine{
	//	Scheduler:   &scheduler.SimpleScheduler{},
	//	WorkerCount: 10,
	//}
	//e1.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: citylist.ParseCityList,
	//})

	//-------------------------------------
	//e2 := engine.ConcurrentEngine{
	//	Scheduler:   &scheduler.QueuedScheduler{},
	//	WorkerCount: 10,
	//}
	//e2.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: citylist.ParseCity,
	//})

	//itemChan, err := persist.ItemSaver("dating_profile")
	//if err != nil {
	//	panic(err)
	//}
	//e3 := engine.ConcurrentEngine{
	//	Scheduler:        &scheduler.QueuedScheduler{},
	//	WorkerCount:      10,
	//	ItemChan:         itemChan,
	//	RequestProcessor: engine.Worker,
	//}
	//e3.Run(engine.Request{
	//	Url:    "http://www.zhenai.com/zhenghun",
	//	Parser: engine.NewFuncParser(citylist.ParseCityList, config.ParseCityList),
	//})

	itemChan, err := persist.ItemSaver("douban_house_online")
	if err != nil {
		panic(err)
	}
	e3 := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      10,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}
	e3.Run(engine.Request{
		Url:    "https://www.douban.com/group/search?cat=1019&q=%E5%8C%97%E4%BA%AC%E7%A7%9F%E6%88%BF",
		Parser: engine.NewFuncParser(douban.ParseSearch, config.ParseCityList),
	})
}
