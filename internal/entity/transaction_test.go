package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTransaction(t *testing.T) {
	clientFrom, err := NewClient("John Doe", "john@doe.com")
	accountFrom := NewAccount(clientFrom)

	clientTo, err := NewClient("Maria Doe", "maria@doe.com")
	accountTo := NewAccount(clientTo)

	accountFrom.Credit(1000)
	accountTo.Credit(1000)

	transaction, err := NewTransaction(accountFrom, accountTo, 100)

	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, 900.0, accountFrom.Balance)
	assert.Equal(t, 1100.0, accountTo.Balance)
}

func TestCreateTransactionWithInsufficientBalance(t *testing.T) {
	clientFrom, err := NewClient("John Doe", "john@doe.com")
	accountFrom := NewAccount(clientFrom)

	clientTo, err := NewClient("Maria Doe", "maria@doe.com")
	accountTo := NewAccount(clientTo)

	accountFrom.Credit(1000)
	accountTo.Credit(1000)

	transaction, err := NewTransaction(accountFrom, accountTo, 2000)
	assert.NotNil(t, err)
	assert.Error(t, err, "insufficient balance")
	assert.Nil(t, transaction)
	assert.Equal(t, 1000.0, accountFrom.Balance)
	assert.Equal(t, 1000.0, accountTo.Balance)
}
