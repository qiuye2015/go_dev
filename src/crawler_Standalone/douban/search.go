package xiaozu

import (
	"github.com/qiuye2015/go_dev/crawler_Standalone/engine"
	"regexp"
)

//
var gUrlRe = regexp.MustCompile(`<h3><a class=""href="(https://www.douban.com/group/.*/)"[^>]*>([^<]+)</a>`)

func ParseSearch(contents []byte) engine.ParseResult {
	result := engine.ParseResult{}
	//ioutil.WriteFile("test_fjp.html", contents, 0777)
	matches := gUrlRe.FindAllSubmatch(contents, -1)

	for _, m := range matches {
		tmp := string(m[1])
		//fmt.Printf("url: %s\n", tmp)
		result.Requests = append(result.Requests, engine.Request{
			Url: tmp,
			//ParserFunc: engine.NilParser,
			ParserFunc: ParseGroup,
		})
		break //for test
	}

	return result
}
