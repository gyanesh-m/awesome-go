package models

import "errors"

// Accounts struct for gorm
type Accounts struct {
	ID             string `gorm:"primaryKey"`
	Amount         int64  `gorm:"column:amount"`
	DocumentNumber string `gorm:"column:document_number"`
	CreatedAt      *int64 `gorm:"column:created_at"`
	UpdatedAt      *int64 `gorm:"column:updated_at"`
	DeletedAt      *int64 `gorm:"column:deleted_at"`
}

func (a Accounts) CanProceedWithTxnAmount(txnAmount int64) error {
	if a.Amount+txnAmount >= 0 {
		return nil
	} else {
		return errors.New("account balance is not sufficient to proceed with the transaction")
	}
}
