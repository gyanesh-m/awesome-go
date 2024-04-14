package cardholder

import (
	"awesome-go/internal/boot"
	v1 "awesome-go/rpc/protos"
	"context"
)

type cardholderServer struct {
	CardHolderService ServiceI
	v1.UnimplementedCardHolderServiceServer
}

// new cs
func NewCardHolderServer(service ServiceI) v1.CardHolderServiceServer {
	return &cardholderServer{
		CardHolderService: service,
	}

}

func (cs *cardholderServer) CreateAccount(ctx context.Context, request *v1.CreateAccountRequest) (*v1.CreateAccountResponse, error) {
	boot.Lgr.Infow("CreateAccount", "request", request)
	model, err := cs.CardHolderService.CreateAccount(ctx, request)
	if err != nil {
		return nil, err
	}
	response := CreateAccountResponse(model)
	return response, nil
}

func (cs *cardholderServer) GetAccount(ctx context.Context, request *v1.GetAccountRequest) (*v1.GetAccountResponse, error) {
	boot.Lgr.Infow("GetAccount", "request", request)
	model, err := cs.CardHolderService.GetAccount(ctx, request)
	if err != nil {
		return nil, err
	}
	response := GetAccountResponse(model)
	return response, nil
}
