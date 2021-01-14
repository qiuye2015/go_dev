package xiaozu

import (
	"crawler_Standalone/engine"
	"fmt"
	"regexp"
	"strings"
)

//`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
//<h1 class="title">
//惠新西街附近主卧招租1500（限女生）
//</h1>
//<h1>
//4号线地铁沿线，精装修，无中介费，拎包入住
//</h1>
//<span class="from">来自: <a href="https://www.douban.com/people/227608756/">Yi.</a></span>
//<span class="create-time color-green" style="display:inline-block">2020-12-15 10:12:18</span>
var titleRe = regexp.MustCompile(`<h1[^>]*>([^<]*)</h1>`)

//var nameRe = regexp.MustCompile(`<span[^>]*>([^<]*)</span>`)
var nameRe = regexp.MustCompile(`<span class="from">[^>]+>([^<]*)</a></span>`)

//var timeRe = regexp.MustCompile(`<span[^>]*>([^<]*)</span>`)
var timeRe = regexp.MustCompile(`<span class="create-time[^>]+>([^<]*)</span>`)

var detailRe = regexp.MustCompile(`<p data-align="left">([^<]*)</p>`)

func ParseHouse(contents []byte) engine.ParseResult {
	result := engine.ParseResult{}
	title := ""
	name := ""
	tm := ""
	detail := ""
	//ioutil.WriteFile("topic_204699328.html", contents, 0777)
	match := titleRe.FindSubmatch(contents)
	//fmt.Println("---------", len(match))
	if len(match) > 1 {
		tmp := strings.TrimLeft(string(match[1]), "\n")
		title = strings.TrimSpace(tmp)
	} else {
		return result
	}

	match = nameRe.FindSubmatch(contents)
	if len(match) > 1 {
		name = string(match[1])
	}
	match = timeRe.FindSubmatch(contents)
	if len(match) > 1 {
		tm = string(match[1])
	}
	matches := detailRe.FindAllSubmatch(contents, -1)

	for _, m := range matches {
		detail += string(m[1])
	}
	fmt.Printf("titile: %s\n", title)
	fmt.Printf("name: %s\n", name)
	fmt.Printf("timestamp: %s\n", tm)
	fmt.Printf("detail: %s\n", detail)

	result.Requests = append(result.Requests, engine.Request{

		ParserFunc: engine.NilParser,
	})

	return result
}
