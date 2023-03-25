package gateway

import "github.com/jefersonalmeida/go-wallet/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
