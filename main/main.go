package main

import (
	"github.com/rezaindrag/goelasticsearch/restapi"

	"github.com/labstack/echo/v4"

	"github.com/olivere/elastic/v7"
	"github.com/rezaindrag/goelasticsearch/repository"
)

func main() {
	e := echo.New()

	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"), elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	restapi.NewAPI(e, repository.NewRepository(client, "document"))

	if err := e.Start(":7723"); err != nil {
		panic(err)
	}
}
