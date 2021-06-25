package xiaozu

import (
	"github.com/qiuye2015/go_dev/crawler_Standalone/engine"
	"regexp"
)

var urlRe1 = regexp.MustCompile(`<a class="item-containor" href="(/group/topic/[0-9]+)" title=*"`)

//<a href="https://www.douban.com/group/topic/205037311/" title="地铁7号线 金蝉南里欢乐谷押一付一 2200次卧合租 百子湾 垡头 化工双合价格超低特别牛特价特价下楼地铁垡" class="">
//地铁7号线 金蝉南里欢乐谷押一付一 2200次卧合租 ...
//</a>

var newUrl = regexp.MustCompile(`<a href="(https://www.douban.com/group/topic/[0-9]+/)"[^>]*>([^<]+)</a>`)

const URL_PREFIX = "https://www.douban.com"

func ParseGroup(contents []byte) engine.ParseResult {
	result := engine.ParseResult{}
	//ioutil.WriteFile("test_fjp.html", contents, 0777)
	matches := newUrl.FindAllSubmatch(contents, -1)
	//fmt.Println("-------1", len(matches))
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			//Url: URL_PREFIX + string(m[1]),
			Url: string(m[1]),
			//ParserFunc: engine.NilParser,
			ParserFunc: ParseHouse,
		})
		//break //for test
	}

	return result
}
