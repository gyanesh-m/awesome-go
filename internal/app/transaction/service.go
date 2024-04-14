package transaction

import (
	"awesome-go/internal/app/cardholder"
	"awesome-go/internal/db/models"
	v1 "awesome-go/rpc/protos"
	"context"
)

// transaction service struct
type service struct {
	TransactionRepository RepositoryI
	TransactionFactory    FactoryI
	CardholderService     cardholder.ServiceI
}

// NewService creates a new transaction service
func NewService(tr RepositoryI, tf FactoryI, cardholderService cardholder.ServiceI) ServiceI {
	return &service{
		TransactionRepository: tr,
		TransactionFactory:    tf,
		CardholderService:     cardholderService,
	}
}

func (s service) CreateTransaction(ctx context.Context, request *v1.CreateTransactionRequest) (*models.Transactions, error) {
	// Can take a lock on account_id to handle for simultaneous updates to the account
	account, err := s.CardholderService.GetAccount(ctx, &v1.GetAccountRequest{AccountId: request.AccountId})
	if err != nil {
		return nil, err
	}

	if err = ValidateCreateTransactionRequest(ctx, request); err != nil {
		return nil, err
	}

	if err = account.CanProceedWithTxnAmount(request.GetAmount()); err != nil {
		return nil, err
	}
	accountModel, err := s.TransactionFactory.BuildTransaction(request)
	if err != nil {
		return nil, err
	}
	if err := s.TransactionRepository.SaveTransaction(accountModel); err != nil {
		return nil, err
	}
	return accountModel, nil
}

func (s service) GetTransaction(ctx context.Context, request *v1.GetTransactionRequest) (*models.Transactions, error) {
	// Can take a lock on transaction_id

	if err := ValidateGetTransactionRequest(ctx, request); err != nil {
		return nil, err
	}
	model, err := s.TransactionRepository.FindTransactionById(request.TransactionId)
	if err != nil {
		return nil, err
	}
	return model, nil
}
