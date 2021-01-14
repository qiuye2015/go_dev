package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"
)

//<a href="http://album.zhenai.com/u/1807074256" target="_blank">余生有你</a>
//var (
//	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
//	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/shanghai/[^"]+)"`)
//)
//
//func ParseCity(contents []byte, url string) engine.ParseResult {
//	result := engine.ParseResult{}
//
//	matches := profileRe.FindAllSubmatch(contents, -1)
//
//	for _, m := range matches {
//		result.Requests = append(result.Requests,
//			engine.Request{
//				Url:        string(m[1]),
//				ParserFunc: ProfileParser(string(m[2])),
//			})
//	}
//
//	//下一页
//	//matches = cityUrlRe.FindAllSubmatch(contents, -1)
//	//for _, m := range matches {
//	//	result.Requests = append(result.Requests,
//	//		engine.Request{
//	//			Url:        string(m[1]),
//	//			ParserFunc: ParseCity,
//	//		})
//	//}
//
//	return result
//}

//<a href="http://album.zhenai.com/u/1807074256" target="_blank">余生有你</a>
var cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
var ageRe1 = regexp.MustCompile(`<span class="grayL">年龄：</span>([^<]+)</td>`)
var marriageRe1 = regexp.MustCompile(`<span class="grayL">婚况：</span>([^<]+)</td>`)
var sexRe1 = regexp.MustCompile(`<span class="grayL">性别：</span>([^<]+)</td>`)
var idUrlRe = regexp.MustCompile(`album.zhenai.com/u/([0-9]+)`)

func ParseCity(contents []byte, _ string) engine.ParseResult {
	result := engine.ParseResult{}

	cityMatches := cityRe.FindAllSubmatch(contents, -1)
	ageMatches := ageRe1.FindAllSubmatch(contents, -1)
	marriageMatches := marriageRe1.FindAllSubmatch(contents, -1)
	sexMatches := sexRe1.FindAllSubmatch(contents, -1)
	//fmt.Printf("city:%d | age:%d | marriage:%d | sex:%d\n",
	//	len(cityMatches), len(ageMatches), len(marriageMatches), len(sexMatches))

	for i, m := range cityMatches {
		profile := model.Profile{}
		profile.Name = string(m[2])
		url := string(m[1])
		profile.Url = url
		if age, err := strconv.Atoi(string(ageMatches[i][1])); err == nil {
			profile.Age = age

		}
		profile.Sex = string(sexMatches[i][1])
		profile.Marriage = string(marriageMatches[i][1])
		//fmt.Printf("name:%s | url:%s | age:%d | sex:%s |marriage:%s\n",
		//	profile.Name, profile.Url, profile.Age, profile.Sex, profile.Marriage)
		result.Items = append(result.Items,
			engine.Item{
				Url:     url,
				Id:      extractString([]byte(url), idUrlRe),
				Payload: profile,
			})
	}

	return result
}
