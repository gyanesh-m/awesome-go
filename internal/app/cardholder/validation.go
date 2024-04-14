package cardholder

import (
	"awesome-go/internal/boot"
	v1 "awesome-go/rpc/protos"
	"context"
	_ "github.com/go-ozzo/ozzo-validation/v4"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Validates the CreateAccountRequest
func ValidateCreateAccountRequest(ctx context.Context, request *v1.CreateAccountRequest) error {
	if err := validation.ValidateStruct(request,
		validation.Field(&request.DocumentNumber, validation.Required),
	); err != nil {
		boot.Lgr.Infof("error validating request: %v", err)
		return err
	}
	return nil
}

// Validate the GetAccountRequest
func ValidateGetAccountRequest(ctx context.Context, request *v1.GetAccountRequest) error {
	if err := validation.ValidateStruct(request,
		validation.Field(&request.AccountId, validation.Required),
	); err != nil {
		boot.Lgr.Infof("error validating request: %v", err)
		return err
	}
	return nil
}
