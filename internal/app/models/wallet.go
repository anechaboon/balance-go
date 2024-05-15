package models

// Users represents a user in the system
type Wallet struct {
	ID                   uint    `json:"id,omitempty"`
	Title                string  `json:"title"`
	Balance              float64 `json:"balance" gorm:"type:decimal(11,2) unsinged"`
	DefaultExpenseCateID uint    `json:"default_expense_cate_id"`
	Status               uint    `json:"status" gorm:"type:tinyint unsinged; default:1"`
}

// TableName ..
func (Wallet) TableName() string {
	return "wallet"
}
