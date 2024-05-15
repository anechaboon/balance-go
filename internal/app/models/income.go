package models

// Users represents a user in the system
type Income struct {
	ID               uint    `json:"id,omitempty"`
	IncomeCategoryID uint    `json:"income_category_id"`
	Memo             string  `json:"memo"`
	Amount           float64 `json:"amount" gorm:"type:decimal(11,2) unsinged"`
	WalletID         uint
	TypeName         string `json:"type_name"` // from join to income_categories
	UserID           uint
	Status           uint `json:"status" gorm:"type:tinyint unsinged; default:1"`
}

// TableName ..
func (Income) TableName() string {
	return "incomes"
}
