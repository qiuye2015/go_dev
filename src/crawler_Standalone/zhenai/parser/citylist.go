package parser

import (
	"crawler_Standalone/engine"
	"fmt"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contensts []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	result := engine.ParseResult{}
	mathes := re.FindAllSubmatch(contensts, -1)

	limit := 2
	for _, m := range mathes {
		result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
		limit--
		if limit == 0 {
			break
		}
		//fmt.Printf("ParseCityList----City: %s, URL: %s\n", m[2], m[1])
	}
	fmt.Printf("ParseCityList----get all Items: %d\n", len(result.Items))
	return result
}
