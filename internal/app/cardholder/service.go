package cardholder

import (
	"awesome-go/internal/db/models"
	v1 "awesome-go/rpc/protos"
	"context"
)

// cardholder service struct
type service struct {
	CardholderRepository RepositoryI
	CardholderFactory    FactoryI
}

// NewService creates a new cardholder service
func NewService(cr RepositoryI, cf FactoryI) ServiceI {
	return &service{
		CardholderRepository: cr,
		CardholderFactory:    cf,
	}
}

func (s service) CreateAccount(ctx context.Context, request *v1.CreateAccountRequest) (*models.Accounts, error) {
	if err := ValidateCreateAccountRequest(ctx, request); err != nil {
		return nil, err
	}
	accountModel, err := s.CardholderFactory.BuildAccount(request)
	if err != nil {
		return nil, err
	}
	if err := s.CardholderRepository.SaveAccount(accountModel); err != nil {
		return nil, err
	}
	return accountModel, nil
}

func (s service) GetAccount(ctx context.Context, request *v1.GetAccountRequest) (*models.Accounts, error) {
	if err := ValidateGetAccountRequest(ctx, request); err != nil {
		return nil, err
	}
	model, err := s.CardholderRepository.FindAccountById(request.AccountId)
	if err != nil {
		return nil, err
	}
	return model, nil
}
