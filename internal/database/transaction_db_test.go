package database

import (
	"database/sql"
	"github.com/jefersonalmeida/go-wallet/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	clientFrom    *entity.Client
	clientTo      *entity.Client
	accountFrom   *entity.Account
	accountTo     *entity.Account
	transactionDB *TransactionDB
}

func (s *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)

	s.db = db
	db.Exec("create table clients (id varchar(255), name varchar(255), email varchar(255), created_at date, updated_at date)")
	db.Exec("create table accounts (id varchar(255), client_id varchar(255), balance float, created_at date, updated_at date)")
	db.Exec("create table transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount float, created_at date)")

	// creating clients
	clientFrom, err := entity.NewClient("John Doe", "john@doe.com")
	s.Nil(err)
	s.clientFrom = clientFrom

	clientTo, err := entity.NewClient("Maria Doe", "maria@doe.com")
	s.Nil(err)
	s.clientTo = clientTo

	// creating accounts
	accountFrom := entity.NewAccount(clientFrom)
	accountFrom.Balance = 1000
	s.Nil(err)
	s.accountFrom = accountFrom

	accountTo := entity.NewAccount(clientTo)
	accountTo.Balance = 1000
	s.Nil(err)
	s.accountTo = accountTo

	s.transactionDB = NewTransactionDB(db)
}

func (s *TransactionDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("drop table clients")
	s.db.Exec("drop table accounts")
	s.db.Exec("drop table transactions")
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (s *TransactionDBTestSuite) TestCreate() {
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 100)
	s.Nil(err)

	err = s.transactionDB.Create(transaction)
	s.Nil(err)
}
