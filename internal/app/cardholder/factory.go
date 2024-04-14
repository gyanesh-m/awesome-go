package cardholder

import (
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

func (cf *factory) BuildAccount(input *v1.CreateAccountRequest) (*models.Accounts, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	currentTime := time.Now().Unix()

	return &models.Accounts{
		ID:             id.String(),
		Amount:         0,
		DocumentNumber: input.DocumentNumber,
		CreatedAt:      &currentTime,
		UpdatedAt:      &currentTime,
	}, nil
}
