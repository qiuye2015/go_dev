package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	//contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseCityList(contents)
	const reslutSize = 470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCities := []string{
		"City 阿坝",
		"City 阿克苏",
		"City 阿拉善盟",
	}
	if len(result.Requests) != reslutSize {
		t.Errorf("result should have %d"+
			"requests; but had %d",
			reslutSize, len(result.Requests))
	}
	if len(result.Items) != reslutSize {
		t.Errorf("result should have %d"+
			"Items; but had %d",
			reslutSize, len(result.Items))
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s;but was %s",
				i, url, result.Requests[i].Url)
		}
	}
	for i, city := range expectedCities {
		if result.Items[i] != city {
			t.Errorf("expected city #%d: %s;but was %s",
				i, city, result.Items[i])
		}
	}

}
