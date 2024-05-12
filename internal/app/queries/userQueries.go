package queries

import (
	"balance-go/internal/app/models"
	"errors"
	"fmt"

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

func (uq *UserQueries) GetUser(id int) (models.Users, error) {
	var user models.Users
	result := uq.db.Find(&user, id)

	if err := result.Error; err != nil {
		return user, err
	}

	// Check if user is empty (indicates not found)
	if user == (models.Users{}) {
		return user, errors.New(fmt.Sprintf("user id [%v] not found", id))
	}

	return user, nil
}

// ... other user-related queries
