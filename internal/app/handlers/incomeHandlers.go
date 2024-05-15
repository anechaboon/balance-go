package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"balance-go/internal/app/queries"
)

type IncomeHandler struct {
	incomeQueries *queries.IncomeQueries
}

func NewIncomeHandler(uq *queries.IncomeQueries) *IncomeHandler {
	return &IncomeHandler{incomeQueries: uq}
}

func (uh *IncomeHandler) GetAllIncomes(c echo.Context) error {
	incomes, err := uh.incomeQueries.GetAllIncomes()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, incomes)
}

func (uh *IncomeHandler) GetIncome(c echo.Context) error {
	// string to int
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		panic(errConv)
	}

	income, err := uh.incomeQueries.GetIncome(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, income)
}

func (uh *IncomeHandler) StoreIncome(c echo.Context) error {

	income, err := uh.incomeQueries.StoreIncome(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, income)
}

// ... other income-related handler functions
