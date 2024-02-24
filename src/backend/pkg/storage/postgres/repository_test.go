package postgres

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"log"
	"testing"

	"github.com/thirteenths/WEB_BMSTU23/backend/internal/model"
	postgres_container "github.com/thirteenths/WEB_BMSTU23/backend/pkg/container/postgres-container"
)

type RepoTestSuite struct {
	suite.Suite
	pgContainer *postgres_container.PostgresContainer
	repository  *Repository
	ctx         context.Context
}

type ModelTestSuite struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	About string `json:"about"`
}

func (suite *RepoTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	pgContainer, err := postgres_container.CreatePostgresContainer(suite.ctx)
	if err != nil {
		log.Fatal(err)
	}
	suite.pgContainer = pgContainer
	conn, err := NewConnect(suite.pgContainer.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}

	repository := NewRepository(model.Event{}, *conn)
	suite.repository = repository
}

func (suite *RepoTestSuite) TearDownSuite() {
	if err := suite.pgContainer.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}

func (suite *RepoTestSuite) TestInsertModel() {
	t := suite.T()

	m, err := suite.repository.Insert(model.Event{
		Id:    1,
		Name:  "name",
		About: "about",
	})

	assert.NoError(t, err)
	assert.NotNil(t, m)
}

func TestRepository_Insert(t *testing.T) {
	suite.Run(t, new(RepoTestSuite))
}

func TestRepository_SelectById(t *testing.T) {

}

func TestRepository_SelectByField(t *testing.T) {

}

func TestRepository_SelectAll(t *testing.T) {

}

func TestRepository_UpdateById(t *testing.T) {

}

func TestRepository_DeleteById(t *testing.T) {

}
