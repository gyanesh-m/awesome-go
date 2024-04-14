package transaction

import (
	"awesome-go/internal/boot"
	v1 "awesome-go/rpc/protos"
	"context"
)

type transactionServer struct {
	TransactionService ServiceI
	v1.UnimplementedTransactionServiceServer
}

// new transaction server
func NewTransactionServer(service ServiceI) v1.TransactionServiceServer {
	return &transactionServer{
		TransactionService: service,
	}
}

// CreateTransaction creates a new transaction for an account_id
func (ts *transactionServer) CreateTransaction(ctx context.Context, request *v1.CreateTransactionRequest) (*v1.CreateTransactionResponse, error) {
	boot.Lgr.Infow("CreateTransaction", "request", request)
	model, err := ts.TransactionService.CreateTransaction(ctx, request)
	if err != nil {
		return nil, err
	}
	response := CreateTransactionResponse(model)
	return response, nil
}

// GetTransaction returns a transaction
func (ts *transactionServer) GetTransaction(ctx context.Context, request *v1.GetTransactionRequest) (*v1.GetTransactionResponse, error) {
	boot.Lgr.Infow("GetTransaction", "request", request)
	model, err := ts.TransactionService.GetTransaction(ctx, request)
	if err != nil {
		return nil, err
	}
	response := GetTransactionResponse(model)
	return response, nil
}
