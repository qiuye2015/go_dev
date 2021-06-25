package parser

import (
	"github.com/qiuye2015/go_dev/crawler/engine"
	"github.com/qiuye2015/go_dev/crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseCity(t *testing.T) {
	contents, err := ioutil.ReadFile("city_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseCity(contents, "")

	//if len(result.Requests) == 0 {
	//	t.Errorf("Error %v", result.Requests)
	//}
	//url := "http://album.zhenai.com/u/1807074256"
	////t.Logf("%+v", result)
	//if result.Requests[0].Url != url {
	//	t.Errorf("result expectd %v; bug was %v", url, result.Requests[0])
	//}

	profile := result.Items[0]
	url := "http://album.zhenai.com/u/1807074256"
	expected := engine.Item{
		Url: url,
		Id:  "1807074256",
		Payload: model.Profile{
			Age:      44,
			Name:     "余生有你",
			Marriage: "离异",
			Sex:      "女士",
			Url:      url,
		},
	}
	if profile != expected {
		t.Errorf("expected %v; but was %v", expected, profile)
	}
}
