package douban

import (
	"fmt"
	"github.com/qiuye2015/go_dev/crawler/engine"
	"regexp"
)

var newUrl = regexp.MustCompile(`<a href="(https://www.douban.com/group/topic/[0-9]+/)"[^>]*>([^<]+)</a>`)

func ParseGroup(contents []byte, _ string) engine.ParseResult {
	result := engine.ParseResult{}
	//ioutil.WriteFile("test_fjp.html", contents, 0777)
	matches := newUrl.FindAllSubmatch(contents, -1)

	//limit := 7
	for _, m := range matches {
		url := string(m[1])
		//if limit == 1 {
		fmt.Printf("url: %s\n", url)
		result.Requests = append(result.Requests, engine.Request{
			Url:    url,
			Parser: engine.NewFuncParser(ParseHouse, url),
		})
		//}
		//limit--
		//if limit == 0 {
		//	break
		//}
	}

	return result
}
