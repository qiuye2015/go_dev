package parser

import (
	"github.com/qiuye2015/go_dev/crawler_Standalone/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseProfile(contents, "余生有你")
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 "+
			"element; but was %v", result.Items)
	}
	profile := result.Items[0].(model.Profile)
	expected := model.Profile{
		Age:      44,
		Name:     "余生有你",
		Marriage: "离异",
	}
	if profile != expected {
		t.Errorf("expected %v; but was %v", expected, profile)
	}
}
