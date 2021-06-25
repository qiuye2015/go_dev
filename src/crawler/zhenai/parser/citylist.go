package parser

import (
	"github.com/qiuye2015/go_dev/crawler/config"
	"github.com/qiuye2015/go_dev/crawler/engine"
	"regexp"
)

//<a data-v-1573aa7c="" href="http://www.zhenai.com/zhenghun/dalian">大连</a>
//<a href="http://www.zhenai.com/zhenghun/dalian" data-v-1573aa7c>大连</a>
//re, err := regexp.Compile(`<a href="http://www.zhenai.com/zhenghun/[0-9a-z]+"[^>]*>[^<]+</a>`)
const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contensts []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	result := engine.ParseResult{}
	mathes := re.FindAllSubmatch(contensts, -1)

	limit := 2
	for _, m := range mathes {
		//fmt.Printf("City: %s, URL: %s\n", m[2], m[1])
		//result.Items = append(result.Items, "City "+string(m[2]))
		url := string(m[1])
		result.Requests = append(result.Requests, engine.Request{
			Url:    url,
			Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
		})
		limit--
		if limit == 0 {
			break
		}

	}
	return result
}
