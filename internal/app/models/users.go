package models

import "time"

// Users ..
type Users struct {
	ID              uint `gorm:"primaryKey" json:"id"`
	Email           string
	Password        string
	FullName        string `json:"full_name"`
	DefaultWalletID uint
	Income          float64 `gorm:"type:decimal(11,2) unsinged"`
	Expense         float64 `gorm:"type:decimal(11,2) unsinged"`
	Balance         float64 `gorm:"type:decimal(11,2) unsinged"`
	ImageURL        string
	RememberToken   string
	Status          uint
	CreatedDate     time.Time `gorm:"type:timestamp;not null;" json:"created_date"`
	UpdatedDate     time.Time `gorm:"type:timestamp;not null;" json:"updated_date"`
}

// TableName ..
func (Users) TableName() string {
	return "users"
}
