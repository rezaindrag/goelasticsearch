package repository

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/olivere/elastic"
	"github.com/stretchr/testify/suite"
)

type repositorySuite struct {
	suite.Suite
	client    *elastic.Client
	indexName string
}

func (r repositorySuite) SetupSuite() {
	fmt.Println("SetupSuite")

}

func (r repositorySuite) TearDownSuite() {
	fmt.Println("TearDownSuite")

}

func (r repositorySuite) SetupTest() {
	fmt.Println("SetupTest")

}

func (r repositorySuite) TearDownTest() {
	fmt.Println("TearDownTest")
}

func TestRepositorySuite(t *testing.T) {
	suite.Run(t, &repositorySuite{})
}

type repositoryTest struct {
	repositorySuite
}

func (r repositoryTest) TestRepository_Fetch() {
	assert.Equal(r.T(), 1, 1)

	fmt.Println("Fetch")
}

func (r repositoryTest) TestRepository_GetByID() {
	assert.Equal(r.T(), 1, 1)

	fmt.Println("GetByID")
}

func (r repositoryTest) TestRepository_Store() {
	assert.Equal(r.T(), 1, 1)

	fmt.Println("Store")
}

func (r repositoryTest) TestRepository_Update() {
	assert.Equal(r.T(), 1, 1)

	fmt.Println("Update")
}

func (r repositoryTest) TestRepository_Delete() {
	assert.Equal(r.T(), 1, 1)

	fmt.Println("Delete")
}
