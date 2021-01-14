package parser

import (
	"crawler_Standalone/engine"
	"crawler_Standalone/model"
	"regexp"
	"strconv"
)

var marriageRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([^<]+)</div>`)
var ageRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([^<]+)岁</div>`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name

	profile.Marriage = extractString(contents, marriageRe)
	if age, err := strconv.Atoi(extractString(contents, ageRe)); err == nil {
		profile.Age = age
	}

	result := engine.ParseResult{
		Requests: nil,
		Items:    []interface{}{profile},
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
