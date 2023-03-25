package database

import (
	"database/sql"
	"github.com/jefersonalmeida/go-wallet/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDB *AccountDB
	client    *entity.Client
}

func (s *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)

	s.db = db

	_, err = db.Exec("create table clients (id varchar(255), name varchar(255), email varchar(255), created_at date, updated_at date)")
	if err != nil {
		return
	}

	_, err = db.Exec("create table accounts (id varchar(255), client_id varchar(255), balance float, created_at date, updated_at date)")
	if err != nil {
		return
	}
	s.accountDB = NewAccountDB(db)

	s.client, _ = entity.NewClient("John Doe", "john@doe.com")
}

func (s *AccountDBTestSuite) TearDownSuite() {
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(s.db)

	_, err := s.db.Exec("drop table clients")
	if err != nil {
		return
	}

	_, err = s.db.Exec("drop table accounts")
	if err != nil {
		return
	}
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (s *AccountDBTestSuite) TestSave() {
	account := entity.NewAccount(s.client)
	err := s.accountDB.Save(account)

	s.Nil(err)
}

func (s *AccountDBTestSuite) TestGet() {

	s.db.Exec("insert into clients (id, name, email, created_at, updated_at) values (?, ?, ?, ?, ?)",
		s.client.ID, s.client.Name, s.client.Email, s.client.CreatedAt, s.client.UpdatedAt,
	)

	account := entity.NewAccount(s.client)
	err := s.accountDB.Save(account)

	s.Nil(err)

	accountDB, err := s.accountDB.Get(account.ID)
	s.Nil(err)
	s.Equal(account.ID, accountDB.ID)
	//s.Equal(account.Client, accountDB.Client)
	s.Equal(account.Client.ID, accountDB.Client.ID)
	s.Equal(account.Client.Name, accountDB.Client.Name)
	s.Equal(account.Client.Email, accountDB.Client.Email)

	s.Equal(account.Balance, accountDB.Balance)
	s.True(account.CreatedAt.Equal(accountDB.CreatedAt))
	s.True(account.UpdatedAt.Equal(accountDB.UpdatedAt))
}
