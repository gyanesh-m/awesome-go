package models

// Cardholder struct for gorm
type Cardholder struct {
	ID        string `gorm:"primaryKey"`
	Name      string `gorm:"column:name"`
	CreatedAt *int64 `gorm:"column:created_at"`
	UpdatedAt *int64 `gorm:"column:updated_at"`
	DeletedAt *int64 `gorm:"column:deleted_at"`
}
