package queries

import (
	"balance-go/internal/app/models"

	"gorm.io/gorm"
)

type UserQueries struct {
	db *gorm.DB
}

func NewUserQueries(db *gorm.DB) *UserQueries {
	return &UserQueries{db: db}
}

func (uq *UserQueries) GetAllUsers() ([]models.Users, error) {
	var users []models.Users
	result := uq.db.Find(&users)
	return users, result.Error
}

// ... other user-related queries
