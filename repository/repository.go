package repository

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/olivere/elastic/v7"
	"github.com/rezaindrag/goelasticsearch"
)

type repository struct {
	client    *elastic.Client
	indexName string
}

func (r repository) Fetch(ctx context.Context) ([]map[string]interface{}, error) {
	// GET /index-name/_search
	result, err := r.client.Search().Index(r.indexName).Do(ctx)
	if err != nil {
		return nil, err
	}

	documents := make([]map[string]interface{}, 0)
	for _, hit := range result.Hits.Hits {
		var document map[string]interface{}
		if err := json.Unmarshal(hit.Source, &document); err != nil {
			log.Println("error marshalling document:", err)
		}
		documents = append(documents, document)
	}

	return documents, nil
}

func (r repository) Update(ctx context.Context, id string, document map[string]interface{}) error {
	// POST /index-name/_doc/1
	_, err := r.client.Index().Index(r.indexName).Id(id).BodyJson(document).Do(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r repository) Delete(ctx context.Context, id string) error {
	// DELETE /index-name/_doc/1
	_, err := r.client.Delete().Index(r.indexName).Id(id).Do(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r repository) GetByID(ctx context.Context, id string) (map[string]interface{}, error) {
	// GET /index-name/_doc/1
	result, err := r.client.Get().Index(r.indexName).Id(id).Do(ctx)
	if err != nil {
		return nil, err
	}

	source, err := result.Source.MarshalJSON()
	if err != nil {
		return nil, err
	}

	document := make(map[string]interface{})
	if err = json.Unmarshal(source, &document); err != nil {
		return nil, err
	}

	return document, nil
}

func (r repository) Store(ctx context.Context, document map[string]interface{}) error {
	id, ok := document["id"].(string)
	if !ok || id == "" {
		return errors.New("id is not provided")
	}

	// POST /index-name/_doc/1
	_, err := r.client.Index().Index(r.indexName).Id(id).BodyJson(document).Do(ctx)
	if err != nil {
		return err
	}

	return nil
}

func NewRepository(client *elastic.Client, indexName string) goelasticsearch.Repository {
	return repository{client: client, indexName: indexName}
}
