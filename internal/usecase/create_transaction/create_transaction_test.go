package create_transaction

import (
	"github.com/jefersonalmeida/go-wallet/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) Get(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	clientFrom, _ := entity.NewClient("John Doe", "john@doe.com")
	accountFrom := entity.NewAccount(clientFrom)
	accountFrom.Credit(1000)

	clientTo, _ := entity.NewClient("Maria Doe", "maria@doe.com")
	accountTo := entity.NewAccount(clientTo)
	accountTo.Credit(1000)

	accountMock := &AccountGatewayMock{}
	accountMock.On("Get", accountFrom.ID).Return(accountFrom, nil)
	accountMock.On("Get", accountTo.ID).Return(accountTo, nil)

	transactionMock := &TransactionGatewayMock{}
	transactionMock.On("Create", mock.Anything).Return(nil)

	uc := NewCreateTransactionUseCase(transactionMock, accountMock)

	inputDTO := CreateTransactionInputDTO{
		AccountIDFrom: accountFrom.ID,
		AccountIDTo:   accountTo.ID,
		Amount:        200,
	}

	output, err := uc.Execute(inputDTO)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)

	accountMock.AssertExpectations(t)
	accountMock.AssertNumberOfCalls(t, "Get", 2)

	transactionMock.AssertExpectations(t)
	transactionMock.AssertNumberOfCalls(t, "Create", 1)
}
