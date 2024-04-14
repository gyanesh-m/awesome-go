package transaction

import (
	"awesome-go/internal/db/models"
	v1 "awesome-go/rpc/protos"
	"context"
)

type ServiceI interface {
	CreateTransaction(ctx context.Context, request *v1.CreateTransactionRequest) (*models.Transactions, error)
	GetTransaction(ctx context.Context, request *v1.GetTransactionRequest) (*models.Transactions, error)
}

type RepositoryI interface {
	SaveTransaction(transactionModel *models.Transactions) error
	FindTransactionById(id string) (*models.Transactions, error)
}

type FactoryI interface {
	BuildTransaction(request *v1.CreateTransactionRequest) (*models.Transactions, error)
}
