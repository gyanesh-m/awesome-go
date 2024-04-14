package models

import "awesome-go/internal/app/enums"

// Transactions struct for gorm
type Transactions struct {
	ID        string                 `gorm:"primaryKey"`
	AccountID string                 `gorm:"column:account_id"`
	Type      *enums.TransactionType `gorm:"column:type"`
	Amount    int64                  `gorm:"column:amount"`
	EventTs   int64                  `gorm:"column:event_ts"`
	CreatedAt *int64                 `gorm:"column:created_at"`
	UpdatedAt *int64                 `gorm:"column:updated_at"`
	DeletedAt *int64                 `gorm:"column:deleted_at"`
}
