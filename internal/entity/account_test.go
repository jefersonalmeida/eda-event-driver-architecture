package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNewAccount(t *testing.T) {
	client, err := NewClient("John Doe", "john@doe.com")
	account := NewAccount(client)

	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.NotNil(t, account)
	assert.Equal(t, client, account.Client)
	assert.Equal(t, 0.0, account.Balance)
}

func TestCreateNewAccountWithNilClient(t *testing.T) {
	account := NewAccount(nil)

	assert.Nil(t, account)
}

func TestCreateAccountCredit(t *testing.T) {
	client, _ := NewClient("John Doe", "john@doe.com")
	account := NewAccount(client)
	account.Credit(100.0)

	assert.Equal(t, 100.0, account.Balance)
}

func TestCreateAccountDebit(t *testing.T) {
	client, _ := NewClient("John Doe", "john@doe.com")
	account := NewAccount(client)
	account.Credit(100.0)
	account.Debit(50.0)

	assert.Equal(t, 50.0, account.Balance)
}
