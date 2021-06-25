package persist

import (
	"github.com/olivere/elastic"
	"github.com/qiuye2015/go_dev/crawler/engine"
	"golang.org/x/net/context"
	"log"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetSniff(false),
		elastic.SetURL("http://10.22.5.25:9200/"),
	)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: Got item "+"#%d: %v", itemCount, item)
			itemCount++
			err := Save(client, index, item)
			if err != nil {
				log.Printf("Item Saver: error: saving item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}

func Save(client *elastic.Client, index string, item engine.Item) error {
	indexService := client.Index().
		Index(index).BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.Do(context.Background())
	if err != nil {
		log.Printf("ES save err: %v", err)
		return err
		//} else {
		//	log.Printf("ES save success, ID: %v", item.Id)
	}
	return nil

}
