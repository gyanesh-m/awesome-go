package transaction

import (
	"awesome-go/internal/boot"
	"awesome-go/internal/db/models"
)

type Repository struct {
}

func NewRepository() RepositoryI {
	return &Repository{}
}

func (cr Repository) SaveTransaction(transactionModel *models.Transactions) (err error) {
	db := boot.Db.Create(transactionModel)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (cr Repository) FindTransactionById(id string) (*models.Transactions, error) {
	var account models.Transactions
	db := boot.Db.Where("id", id).
		Where("deleted_at", nil).
		First(&account)
	if db.Error != nil {
		return nil, db.Error
	}
	return &account, nil
}
