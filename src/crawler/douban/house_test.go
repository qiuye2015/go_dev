package douban

import (
	"github.com/qiuye2015/go_dev/crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseHouse(t *testing.T) {
	contents, err := ioutil.ReadFile("House_test_data.html")
	if err != nil {
		panic(err)
	}
	url := "https://www.douban.com/group/topic/205125512/"

	result := ParseHouse(contents, url)

	houseItem := result.Items[0].Payload
	expected := model.HouseItem{
		"巧巧",
		"地铁6号线黄渠站附近，常营民族家园，找人合租，大...",
		"2020-12-18 13:20:44",
		url,
		"",
	}
	if houseItem != expected {
		t.Errorf("expected %v; but was %v", expected, houseItem)
	}
}
