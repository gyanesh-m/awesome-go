package cardholder

import (
	"awesome-go/internal/boot"
	"awesome-go/internal/db/models"
	"errors"
	"gorm.io/gorm"
)

type Repository struct {
}

func NewRepository() RepositoryI {
	return &Repository{}
}
func (cr Repository) SaveAccount(account *models.Accounts) (err error) {
	db := boot.Db.Create(account)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (cr Repository) FindAccountById(id string) (*models.Accounts, error) {
	var account models.Accounts
	db := boot.Db.Where("id", id).
		Where("deleted_at", nil).
		First(&account)
	if db.Error != nil {
		if gorm.ErrRecordNotFound.Error() == db.Error.Error() {
			return nil, errors.New("invalid account id, record not found")
		}
		return nil, db.Error
	}
	return &account, nil
}
