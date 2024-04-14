package cardholder

import (
	"awesome-go/internal/db/models"
	v1 "awesome-go/rpc/protos"
)

func CreateAccountResponse(model *models.Accounts) *v1.CreateAccountResponse {
	return &v1.CreateAccountResponse{
		AccountId:      model.ID,
		DocumentNumber: model.DocumentNumber,
	}
}

func GetAccountResponse(model *models.Accounts) *v1.GetAccountResponse {
	return &v1.GetAccountResponse{
		AccountId:      model.ID,
		DocumentNumber: model.DocumentNumber,
	}
}
