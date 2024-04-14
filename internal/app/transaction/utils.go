package transaction

import (
	"awesome-go/internal/db/models"
	v1 "awesome-go/rpc/protos"
)

func CreateTransactionResponse(model *models.Transactions) *v1.CreateTransactionResponse {
	return &v1.CreateTransactionResponse{
		TransactionId:   model.ID,
		AccountId:       model.AccountID,
		TransactionType: model.Type.Int(),
		Amount:          &model.Amount,
		EventTimestamp:  model.EventTs,
	}
}

func GetTransactionResponse(model *models.Transactions) *v1.GetTransactionResponse {
	return &v1.GetTransactionResponse{
		TransactionId:   model.ID,
		AccountId:       model.AccountID,
		TransactionType: model.Type.Int(),
		Amount:          &model.Amount,
		EventTimestamp:  model.EventTs,
	}
}
