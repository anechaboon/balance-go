package queries

import (
	"balance-go/internal/app/models"
	"errors"
	"fmt"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type IncomeCategoiresQueries struct {
	db *gorm.DB
}

func NewIncomeCategoiresQueries(db *gorm.DB) *IncomeCategoiresQueries {
	return &IncomeCategoiresQueries{db: db}
}

func (icq *IncomeCategoiresQueries) GetAllIncomeCategoires() ([]models.IncomeCategories, error) {
	var incomeCategoires []models.IncomeCategories
	result := icq.db.Find(&incomeCategoires)
	return incomeCategoires, result.Error
}

func (icq *IncomeCategoiresQueries) GetIncomeCategoires(id int) (models.IncomeCategories, error) {
	var incomeCategoires models.IncomeCategories
	result := icq.db.Find(&incomeCategoires, id)

	if err := result.Error; err != nil {
		return incomeCategoires, err
	}

	// Check if incomeCategoires is empty (indicates not found)
	if incomeCategoires == (models.IncomeCategories{}) {
		return incomeCategoires, errors.New(fmt.Sprintf("incomeCategoires id [%v] not found", id))
	}

	return incomeCategoires, nil
}

func (icq *IncomeCategoiresQueries) StoreIncomeCategoires(c echo.Context) (models.IncomeCategories, error) {
	var incomeCategoires models.IncomeCategories
	if err := c.Bind(&incomeCategoires); err != nil {
		return incomeCategoires, err
	}
	result := icq.db.Create(&incomeCategoires)
	if err := result.Error; err != nil {
		return incomeCategoires, err
	}

	return incomeCategoires, nil
}
