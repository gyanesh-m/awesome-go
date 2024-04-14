package transaction

import (
	"awesome-go/internal/app/cardholder"
	"awesome-go/internal/app/enums"
	"awesome-go/internal/boot"
	"awesome-go/internal/db/models"
	cardholder_mocks "awesome-go/internal/mocks/cardholder"
	transaction_mocks "awesome-go/internal/mocks/transaction"
	v1 "awesome-go/rpc/protos"
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TransactionServerTestSuite struct {
	TransactionService ServiceI
	Server             v1.TransactionServiceServer
	Repository         *transaction_mocks.RepositoryI
	Factory            *transaction_mocks.FactoryI
	CardholderService  cardholder.ServiceI
	suite.Suite
}

func (suite *TransactionServerTestSuite) SetupTest() {
	suite.Repository = transaction_mocks.NewRepositoryI(suite.T())
	suite.Factory = transaction_mocks.NewFactoryI(suite.T())
	suite.CardholderService = cardholder_mocks.NewServiceI(suite.T())
	suite.TransactionService = NewService(suite.Repository, suite.Factory, suite.CardholderService)
	suite.Server = NewTransactionServer(suite.TransactionService)
}

func TestTransactionServerTest(t *testing.T) {
	boot.InitLogger()
	suite.Run(t, new(TransactionServerTestSuite))
}

// TestCreateTransaction test the CreateTransaction method
func (suite *TransactionServerTestSuite) TestCreateTransaction() {
	transactionType := int64(1)
	transactionAmount := int64(1000)
	normalPurchase := enums.NormalPurchase
	testData := []struct {
		name          string
		input         *v1.CreateTransactionRequest
		expected      *v1.CreateTransactionResponse
		expectedError error
	}{
		{
			"success",
			&v1.CreateTransactionRequest{
				AccountId:       "123",
				TransactionType: &transactionType,
				Amount:          &transactionAmount,
			},
			&v1.CreateTransactionResponse{
				TransactionId:   "transaction_id_123",
				AccountId:       "123",
				TransactionType: &transactionType,
				Amount:          &transactionAmount,
			},
			nil,
		},
		{
			"error",
			&v1.CreateTransactionRequest{
				AccountId: "123",
				Amount:    &transactionAmount,
			},
			nil,
			errors.New("account not found"),
		},
	}

	for _, tt := range testData {
		suite.Run(tt.name, func() {
			suite.CardholderService.(*cardholder_mocks.ServiceI).On("GetAccount", mock.Anything, mock.Anything).Return(&models.Accounts{
				ID:             "123",
				Amount:         123,
				DocumentNumber: "doc_123",
			}, nil)
			if tt.expectedError != nil {
				suite.CardholderService.(*cardholder_mocks.ServiceI).On("GetAccount", mock.Anything, mock.Anything).Return(nil,
					errors.New("account not found"))
			} else {
				suite.Factory.On("BuildTransaction", mock.Anything).Return(&models.Transactions{
					ID:        "transaction_id_123",
					AccountID: "123",
					Amount:    1000,
					Type:      &normalPurchase,
				}, nil)
				suite.Repository.On("SaveTransaction", mock.Anything).Return(nil)
			}
			response, err := suite.Server.CreateTransaction(context.Background(), tt.input)
			if tt.expectedError != nil {
				suite.Assert().Error(err)
			} else {
				suite.Assert().NoError(err)
				suite.Assert().Equal(tt.expected, response)
			}
		})
	}

}
