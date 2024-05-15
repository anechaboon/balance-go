package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"balance-go/internal/app/queries"
)

type WalletHandler struct {
	walletQueries *queries.WalletQueries
}

func NewWalletHandler(uq *queries.WalletQueries) *WalletHandler {
	return &WalletHandler{walletQueries: uq}
}

func (uh *WalletHandler) GetAllWallets(c echo.Context) error {
	wallets, err := uh.walletQueries.GetAllWallets()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, wallets)
}

func (uh *WalletHandler) GetWallet(c echo.Context) error {
	// string to int
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		panic(errConv)
	}

	wallet, err := uh.walletQueries.GetWallet(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, wallet)
}

func (uh *WalletHandler) StoreWallet(c echo.Context) error {

	wallet, err := uh.walletQueries.StoreWallet(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, wallet)
}

// ... other wallet-related handler functions
