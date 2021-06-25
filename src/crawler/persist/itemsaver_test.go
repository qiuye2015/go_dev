package persist

import (
	"encoding/json"
	"github.com/olivere/elastic"
	"github.com/qiuye2015/go_dev/crawler/engine"
	"github.com/qiuye2015/go_dev/crawler/model"
	"golang.org/x/net/context"
	"testing"
)

func Test_save(t *testing.T) {
	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetSniff(false),
		elastic.SetURL("http://10.22.5.25:9200/"),
	)
	if err != nil {
		panic(err)
	}

	expected := engine.Item{
		Url: "",
		Id:  "12345",
		Payload: model.Profile{
			Age:      44,
			Name:     "余生有你",
			Marriage: "离异",
		},
	}
	const index = "dating_profile"
	err = Save(client, index, expected)
	if err != nil {
		panic(err)
	}
	resp, err := client.Get().Index(index).Id(expected.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}

	t.Logf("%s", resp.Source)
	var actual engine.Item
	json.Unmarshal(resp.Source, &actual)
	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile
	if actual != expected {
		t.Errorf("got %v; expected: %v", actual, expected)
	}

}
