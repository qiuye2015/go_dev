package parser

import (
	"crawler_Standalone/model"
	"io/ioutil"
	"testing"
)

func TestParseCity(t *testing.T) {
	contents, err := ioutil.ReadFile("city_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseCity(contents)

	//expected := []string{
	//	"User 余生有你",
	//}
	//if len(result.Items) == 0 {
	//	t.Errorf("Error %v", result.Items)
	//}
	//for i, name := range expected {
	//	if result.Items[i] != name {
	//		t.Errorf("result expectd %v; bug was %v", name, result.Items[i])
	//	}
	//}
	profile := result.Items[0].(model.Profile)
	expected := model.Profile{
		Age:      44,
		Name:     "余生有你",
		Marriage: "离异",
		Sex:      "女士",
		Url:      "http://album.zhenai.com/u/1807074256",
	}
	if profile != expected {
		t.Errorf("expected %v; but was %v", expected, profile)
	}
}
