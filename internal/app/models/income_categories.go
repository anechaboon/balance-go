package models

// Users represents a user in the system
type IncomeCategories struct {
	ID     uint   `json:"id,omitempty"`
	Title  string `json:"title"`
	Name   string `json:"name"`
	Icon   string `json:"icon"`
	UserID uint
	Status uint `json:"status" gorm:"type:tinyint unsinged; default:1"`
}

// TableName ..
func (IncomeCategories) TableName() string {
	return "income_categories"
}
