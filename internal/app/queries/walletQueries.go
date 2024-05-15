package queries

import (
	"balance-go/internal/app/models"
	"errors"
	"fmt"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type WalletQueries struct {
	db *gorm.DB
}

func NewWalletQueries(db *gorm.DB) *WalletQueries {
	return &WalletQueries{db: db}
}

func (wq *WalletQueries) GetAllWallets() ([]models.Wallet, error) {
	var wallets []models.Wallet
	result := wq.db.Find(&wallets)
	return wallets, result.Error
}

func (wq *WalletQueries) GetWallet(id int) (models.Wallet, error) {
	var wallet models.Wallet
	result := wq.db.Find(&wallet, id)

	if err := result.Error; err != nil {
		return wallet, err
	}

	// Check if wallet is empty (indicates not found)
	if wallet == (models.Wallet{}) {
		return wallet, errors.New(fmt.Sprintf("wallet id [%v] not found", id))
	}

	return wallet, nil
}

func (wq *WalletQueries) StoreWallet_(c echo.Context) (models.Wallet, error) {
	var wallet models.Wallet

	if err := c.Bind(&wallet); err != nil {
		return wallet, err
	}

	if err := wq.db.Save(&wallet).Error; err != nil {
		return wallet, err
	}

	return wallet, nil
}

func (wq *WalletQueries) StoreWallet(c echo.Context) (models.Wallet, error) {
	var wallet models.Wallet
	if err := c.Bind(&wallet); err != nil {
		return wallet, err
	}
	result := wq.db.Create(&wallet)
	if err := result.Error; err != nil {
		return wallet, err
	}
	if result.RowsAffected == 0 {
		fmt.Println("Warning: No rows affected during wallet creation")
	}
	fmt.Println("result.RowsAffected[%v]", result.RowsAffected)

	return wallet, nil
}
