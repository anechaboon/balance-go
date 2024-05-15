package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"balance-go/internal/app/queries"
)

type IncomeCategoiresHandler struct {
	incomeCategoiresQueries *queries.IncomeCategoiresQueries
}

func NewIncomeCategoiresHandler(uq *queries.IncomeCategoiresQueries) *IncomeCategoiresHandler {
	return &IncomeCategoiresHandler{incomeCategoiresQueries: uq}
}

func (uh *IncomeCategoiresHandler) GetAllIncomeCategoires(c echo.Context) error {
	incomeCategoires, err := uh.incomeCategoiresQueries.GetAllIncomeCategoires()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, incomeCategoires)
}

func (uh *IncomeCategoiresHandler) GetIncomeCategoires(c echo.Context) error {
	// string to int
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		panic(errConv)
	}

	incomeCategoires, err := uh.incomeCategoiresQueries.GetIncomeCategoires(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, incomeCategoires)
}

func (uh *IncomeCategoiresHandler) StoreIncomeCategoires(c echo.Context) error {

	incomeCategoires, err := uh.incomeCategoiresQueries.StoreIncomeCategoires(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, incomeCategoires)
}

// ... other incomeCategoires-related handler functions
