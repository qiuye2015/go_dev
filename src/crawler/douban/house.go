package douban

import (
	"github.com/qiuye2015/go_dev/crawler/engine"
	"github.com/qiuye2015/go_dev/crawler/model"
	"regexp"
	"strings"
	"time"
)

var titleRe = regexp.MustCompile(`<h1[^>]*>([^<]*)</h1>`)
var nameRe = regexp.MustCompile(`<span class="from">[^>]+>([^<]*)</a></span>`)
var timeRe = regexp.MustCompile(`<span class="create-time[^>]+>([^<]*)</span>`)

//var detailRe = regexp.MustCompile(`<p data-align="left">([^<]*)</p>`)
var detailRe = regexp.MustCompile(`<p data-align="left">([^<]*)</p>`)

//<tr><td class="tablelc"></td><td class="tablecc"><strong>标题：</strong>鹿港佳苑，超大次卧，落地窗2800随时看房，方便加微信1576981848</td><td class="tablerc"></td></tr>

var urlIdRe = regexp.MustCompile(`https://www.douban.com/group/topic/([0-9a-z]+)`)
var textRe = regexp.MustCompile(`"text":(.*),`)

const (
	format = "2006-01-02 15:04:05"
)

func ParseHouse(contents []byte, url string) engine.ParseResult {
	result := engine.ParseResult{}

	title := ""
	name := ""
	tm := ""
	detail := ""
	//ioutil.WriteFile("topic_fjp.html", contents, 0777)
	match := extractString(contents, titleRe)
	tmp := strings.Trim(match, "\n")
	title = strings.TrimSpace(tmp)

	name = extractString(contents, nameRe)
	tm = extractString(contents, timeRe)

	matches := detailRe.FindAllSubmatch(contents, -1)
	if len(matches) == 0 {
		matches = textRe.FindAllSubmatch(contents, -1)
	}
	for _, m := range matches {
		detail += string(m[1])
	}
	//15天之内得
	publimeTime, _ := time.Parse(format, tm)
	timeLimit := time.Now().Add(-24 * 15 * time.Hour)
	if publimeTime.Before(timeLimit) {
		return result
	}
	if strings.Contains(detail, "立水桥") && strings.Contains(detail, "房东") {
		//
		//fmt.Printf("titile: %s\n", title)
		//fmt.Printf("name: %s\n", name)
		//fmt.Printf("timestamp: %s\n", tm)
		//fmt.Printf("detail: %s\n", detail)
		//fmt.Printf("url: %s\n", url)

		house := model.HouseItem{}
		house.Author = name
		house.Url = url
		house.Detail = detail
		house.Title = title
		house.Detail = detail
		house.Timestamp = tm

		result.Items = append(result.Items,
			engine.Item{
				Url:     url,
				Id:      extractString([]byte(url), urlIdRe),
				Payload: house,
			})
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return " "
	}
}
