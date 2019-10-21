package persist

import (
	"context"
	"errors"
	"go-crawler/engine"
	"log"

	"gopkg.in/olivere/elastic.v5"
)

// ItemSaver func
func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item saver: got item #%d: %v", itemCount, item)
			itemCount++
			err := Save(client, index, item)
			if err != nil {
				log.Printf("Item saver: error saving item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}

func Save(client *elastic.Client, index string, item engine.Item) error {

	if item.Type == "" {
		return errors.New("must supply Type")
	}

	indexService := client.Index().Index(index).Type(item.Type).BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	resp, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}
	log.Printf("%+v", resp)
	return nil
}
