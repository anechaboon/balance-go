package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"balance-go/internal/app/queries"
)

type UserHandler struct {
	userQueries *queries.UserQueries
}

func NewUserHandler(uq *queries.UserQueries) *UserHandler {
	return &UserHandler{userQueries: uq}
}

func (uh *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := uh.userQueries.GetAllUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

// ... other user-related handler functions
