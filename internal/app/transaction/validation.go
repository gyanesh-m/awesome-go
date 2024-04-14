package transaction

import (
	"awesome-go/internal/app/enums"
	"awesome-go/internal/boot"
	v1 "awesome-go/rpc/protos"
	"context"
	_ "github.com/go-ozzo/ozzo-validation/v4"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Validates the CreateTransactionRequest
func ValidateCreateTransactionRequest(ctx context.Context, request *v1.CreateTransactionRequest) error {
	if err := validation.ValidateStruct(request,
		validation.Field(&request.AccountId, validation.Required),
		validation.Field(&request.Amount, validation.Required),
	); err != nil {
		boot.Lgr.Infof("error validating request: %v", err)
		return err
	}
	if err := enums.IsValidTransactionType(request.GetTransactionType()); err != nil {
		return err
	}
	return nil
}

// Validate the GetTransactionRequest
func ValidateGetTransactionRequest(ctx context.Context, request *v1.GetTransactionRequest) error {
	if err := validation.ValidateStruct(request,
		validation.Field(&request.TransactionId, validation.Required),
	); err != nil {
		boot.Lgr.Infof("error validating request: %v", err)
		return err
	}
	return nil
}
