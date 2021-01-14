package main

import (
	xiaozu "crawler_Standalone/douban"
	"crawler_Standalone/engine"
)

func main() {
	//request := engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: citylist.ParseCityList,
	//}
	//engine.Run(request)

	////豆瓣小组
	//request := engine.Request{
	//	Url:        "https://www.douban.com/group/585926/",
	//	ParserFunc: xiaozu.ParseGroup,
	//}
	//搜索豆瓣小组
	request := engine.Request{
		Url:        "https://www.douban.com/group/search?cat=1019&q=%E5%8C%97%E4%BA%AC%E7%A7%9F%E6%88%BF",
		ParserFunc: xiaozu.ParseSearch,
	}
	////每篇文章
	//request := engine.Request{
	//	Url:        "https://www.douban.com/group/topic/204699328/",
	//	ParserFunc: xiaozu.ParseHouse,
	//}
	engine.Run(request)
}
