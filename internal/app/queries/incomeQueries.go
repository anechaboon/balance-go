package queries

import (
	"balance-go/internal/app/models"
	"errors"
	"fmt"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type IncomeQueries struct {
	db *gorm.DB
}

func NewIncomeQueries(db *gorm.DB) *IncomeQueries {
	return &IncomeQueries{db: db}
}

func (icq *IncomeQueries) GetAllIncomes() ([]models.Income, error) {
	var incomes []models.Income

	result := icq.db.Model(&models.Income{}).
		Select("incomes.*, income_categories.name as type_name").
		Joins("JOIN income_categories ON incomes.income_category_id = income_categories.id").
		Where("incomes.status = ?", 1).
		Where("income_categories.status = ?", 1).
		Find(&incomes)

	if err := result.Error; err != nil {
		fmt.Println(err.Error())
		return incomes, err
	}

	return incomes, nil
}

func (wq *IncomeQueries) GetIncome(id int) (models.Income, error) {
	var income models.Income
	result := wq.db.Find(&income, id)

	if err := result.Error; err != nil {
		return income, err
	}

	// Check if income is empty (indicates not found)
	if income == (models.Income{}) {
		return income, errors.New(fmt.Sprintf("income id [%v] not found", id))
	}

	return income, nil
}

func (wq *IncomeQueries) StoreIncome(c echo.Context) (models.Income, error) {
	var income models.Income
	if err := c.Bind(&income); err != nil {
		return income, err
	}
	result := wq.db.Create(&income)
	if err := result.Error; err != nil {
		return income, err
	}

	return income, nil
}
