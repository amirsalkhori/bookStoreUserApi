package elastic

import (
	"bookStoreUser/logger"
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
)

type Elastic struct {
}

func (el Elastic) getESClient() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200/"), elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	return client, nil
}

func (el Elastic) Insert(index string, obj interface{}) error {
	client, err := el.getESClient()
	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic(err)
	}

	indexName := index
	exists, err := client.IndexExists(indexName).Do(context.Background())
	if err != nil {
		panic(err)
	}

	if !exists {
		createIndex, err := client.CreateIndex(indexName).Do(context.Background())
		if err != nil {
			panic(err)
		}

		if !createIndex.Acknowledged {
			fmt.Println("Is not index...")
		}
	}

	data := obj

	_, err = client.Index().
		Index(indexName).
		// Id(data.ID).
		BodyJson(data).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	logger.Info("Data indexed successfully")

	return nil
}

type Student struct {
	Name         string  `json:"name"`
	Age          int64   `json:"age"`
	AverageScore float64 `json:"average_score"`
}
