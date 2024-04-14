package cardholder

import (
	"awesome-go/internal/db/models"
	v1 "awesome-go/rpc/protos"
	"context"
)

type ServiceI interface {
	CreateAccount(ctx context.Context, request *v1.CreateAccountRequest) (*models.Accounts, error)
	GetAccount(ctx context.Context, request *v1.GetAccountRequest) (*models.Accounts, error)
}

type RepositoryI interface {
	SaveAccount(accountModel *models.Accounts) error
	FindAccountById(id string) (*models.Accounts, error)
}

type FactoryI interface {
	BuildAccount(request *v1.CreateAccountRequest) (*models.Accounts, error)
}
