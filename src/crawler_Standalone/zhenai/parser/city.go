package parser

import (
	"crawler_Standalone/engine"
	"crawler_Standalone/model"
	"regexp"
	"strconv"
)

//<a href="http://album.zhenai.com/u/1807074256" target="_blank">余生有你</a>
var cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
var ageRe1 = regexp.MustCompile(`<span class="grayL">年龄：</span>([^<]+)</td>`)
var marriageRe1 = regexp.MustCompile(`<span class="grayL">婚况：</span>([^<]+)</td>`)
var sexRe1 = regexp.MustCompile(`<span class="grayL">性别：</span>([^<]+)</td>`)

//func ParseCity(contents []byte) engine.ParseResult {
//	result := engine.ParseResult{}
//
//	re := regexp.MustCompile(cityRe)
//	matches := re.FindAllSubmatch(contents, -1)
//
//	for _, m := range matches {
//		result.Items = append(result.Items,
//			"User "+string(m[2]))
//		result.Requests = append(result.Requests, engine.Request{
//			Url: string(m[1]),
//			//为了传递name,使用闭包
//			ParserFunc: func(c []byte) engine.ParseResult {
//				return ParseProfile(c, string(m[2]))
//			},
//		})
//	}
//
//	return result
//}

func ParseCity(contents []byte) engine.ParseResult {
	result := engine.ParseResult{}

	cityMatches := cityRe.FindAllSubmatch(contents, -1)
	ageMatches := ageRe1.FindAllSubmatch(contents, -1)
	marriageMatches := marriageRe1.FindAllSubmatch(contents, -1)
	sexMatches := sexRe1.FindAllSubmatch(contents, -1)

	for i, m := range cityMatches {
		profile := model.Profile{}
		profile.Name = string(m[2])
		profile.Url = string(m[1])
		if age, err := strconv.Atoi(string(ageMatches[i][1])); err == nil {
			profile.Age = age

		}
		profile.Sex = string(sexMatches[i][1])
		profile.Marriage = string(marriageMatches[i][1])
		result.Items = append(result.Items, profile)
		//result.Requests = append(result.Requests, engine.Request{
		//	Url: string(m[1]),
		//	ParserFunc:engine.NilParser,
		//	},
		//)
	}

	return result
}
