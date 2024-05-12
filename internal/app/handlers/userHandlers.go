package handlers

import (
	"net/http"
	"strconv"

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

func (uh *UserHandler) GetUser(c echo.Context) error {
	// string to int
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		panic(errConv)
	}

	user, err := uh.userQueries.GetUser(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

// ... other user-related handler functions
