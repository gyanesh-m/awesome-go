package transaction

import (
	"awesome-go/internal/app/enums"
	"awesome-go/internal/db/models"
	v1 "awesome-go/rpc/protos"
	"github.com/google/uuid"
	_ "github.com/google/uuid"
	"time"
)

type factory struct {
}

func NewFactory() FactoryI {
	return &factory{}
}

func (cf *factory) BuildTransaction(request *v1.CreateTransactionRequest) (*models.Transactions, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	currentTime := time.Now().Unix()

	transactionType, err := enums.GetTransactionTypeFromInt(request.GetTransactionType())
	if err != nil {
		return nil, err
	}
	return &models.Transactions{
		ID:        id.String(),
		AccountID: request.GetAccountId(),
		Type:      transactionType,
		Amount:    request.GetAmount() * transactionType.Sign(),
		EventTs:   currentTime,
		CreatedAt: &currentTime,
		UpdatedAt: &currentTime,
	}, nil
}
