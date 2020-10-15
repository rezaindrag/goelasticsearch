package repository

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type repositorySuite struct {
	suite.Suite
	indexName string
	client    *elastic.Client
}

func (r *repositorySuite) SetupSuite() {
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"), elastic.SetSniff(false))
	if err != nil {
		require.NoError(r.T(), err)
	}
	r.client = client
	r.indexName = "document_test"
}

func (r *repositorySuite) SetupTest() {
	result, err := r.client.CreateIndex(r.indexName).Do(context.Background())
	require.NoError(r.T(), err)
	if !result.Acknowledged {
		panic("creating index is unacknowledged")
	}

	// Init mock data
	documents := []map[string]interface{}{
		{
			"id":    "1",
			"title": "foo",
		},
		{
			"id":    "2",
			"title": "bar",
		},
		{
			"id":    "3",
			"title": "blah",
		},
	}
	for _, document := range documents {
		_, err = r.client.Index().Index(r.indexName).Id(uuid.NewV4().String()).BodyJson(document).Do(context.TODO())
		if err != nil {
			require.NoError(r.T(), err)
		}
	}
}

func (r *repositorySuite) TearDownTest() {
	result, err := r.client.DeleteIndex(r.indexName).Do(context.Background())
	require.NoError(r.T(), err)
	if !result.Acknowledged {
		panic("deleting index is unacknowledged")
	}
}

type repositoryTest struct {
	repositorySuite
}

func Test(t *testing.T) {
	suite.Run(t, new(repositoryTest))
}

func (r *repositoryTest) TestRepository_Fetch() {
	assert.Equal(r.T(), "bar", r.indexName)

	repository := NewRepository(r.client, r.indexName)
	documents, err := repository.Fetch(context.TODO())
	require.NoError(r.T(), err)
	assert.Len(r.T(), documents, 3)

	fmt.Println("Fetch")
}

func (r *repositoryTest) TestRepository_GetByID() {
	assert.Equal(r.T(), 1, 1)

	fmt.Println("GetByID")
}

func (r *repositoryTest) TestRepository_Store() {
	assert.Equal(r.T(), 1, 1)

	fmt.Println("Store")
}

func (r *repositoryTest) TestRepository_Update() {
	assert.Equal(r.T(), 1, 1)

	fmt.Println("Update")
}

func (r *repositoryTest) TestRepository_Delete() {
	assert.Equal(r.T(), 1, 1)

	fmt.Println("Delete")
}
