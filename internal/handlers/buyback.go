package handlers

import (
	"errors"

	"github.com/yahfiilham/gold-store-demo/configs"
	"github.com/yahfiilham/gold-store-demo/pkg/models"
)

func SaveBuybackGold(cfg *configs.Config, data *models.BaseRequest) error {
	ac, err := GetAccount(cfg, data.AccountNumber)
	if err != nil {
		return err
	}

	if data.Gold > ac.Balance {
		return errors.New("buyback balance cannot be less than the buyback amount")
	}

	if err := SaveToAccount(cfg, &models.AccountData{
		AccountNumber: data.AccountNumber,
		Balance:       data.Gold,
	}, "buyback"); err != nil {
		return err
	}

	if err := SaveToTransaction(cfg, &models.TransactionData{
		Gold: data.Gold,
		Type: "buyback",
	}, data.AccountNumber); err != nil {
		return err
	}

	return nil
}
