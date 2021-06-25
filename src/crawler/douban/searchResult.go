package douban

import (
	"github.com/qiuye2015/go_dev/crawler/engine"
	"regexp"
)

//
var urlRe = regexp.MustCompile(`<h3><a class=""href="(https://www.douban.com/group/.*/)"[^>]*>([^<]+)</a>`)

func ParseSearch(contents []byte, _ string) engine.ParseResult {
	result := engine.ParseResult{}
	//ioutil.WriteFile("test_fjp.html", contents, 0777)
	matches := urlRe.FindAllSubmatch(contents, -1)

	//limit := 1
	for _, m := range matches {
		url := string(m[1])
		//fmt.Printf("url: %s\n", url)
		result.Requests = append(result.Requests, engine.Request{
			Url:    url,
			Parser: engine.NewFuncParser(ParseGroup, ""),
		})
		//limit--
		//if limit == 0 {
		//	break
		//}
	}

	return result
}
