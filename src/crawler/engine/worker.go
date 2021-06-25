package engine

import (
	"github.com/qiuye2015/go_dev/crawler/fetcher"
	"log"
)

func Worker(r Request) (ParseResult, error) {
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error "+
			"fetching url %s : %v :%v",
			r.Url, err, body)
		return ParseResult{}, err
	}
	return r.Parser.Parse(body, r.Url), nil
}
