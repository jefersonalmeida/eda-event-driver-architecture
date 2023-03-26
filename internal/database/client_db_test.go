package database

import (
	"database/sql"
	"github.com/jefersonalmeida/go-wallet/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ClientDBTestSuite struct {
	suite.Suite
	db       *sql.DB
	clientDB *ClientDB
}

func (s *ClientDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)

	s.db = db
	db.Exec("create table clients (id varchar(255), name varchar(255), email varchar(255), created_at date, updated_at date)")

	s.clientDB = NewClientDB(db)
}

func (s *ClientDBTestSuite) TearDownSuite() {
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(s.db)

	_, err := s.db.Exec("drop table clients")
	if err != nil {
		return
	}
}

func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuite))
}

func (s *ClientDBTestSuite) TestSave() {
	client, _ := entity.NewClient("John Doe", "john@doe.com")
	err := s.clientDB.Save(client)

	s.Nil(err)
}

func (s *ClientDBTestSuite) TestGet() {
	client, _ := entity.NewClient("John Doe", "john@doe.com")
	err := s.clientDB.Save(client)

	s.Nil(err)

	clientDB, err := s.clientDB.Get(client.ID)
	s.Nil(err)
	s.Equal(client.ID, clientDB.ID)
	s.Equal(client.Name, clientDB.Name)
	s.Equal(client.Email, clientDB.Email)
	s.True(client.CreatedAt.Equal(client.CreatedAt))
	s.True(client.UpdatedAt.Equal(client.UpdatedAt))
}
