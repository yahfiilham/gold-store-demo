package handlers

import (
	"errors"

	"github.com/yahfiilham/gold-store-demo/configs"
	"github.com/yahfiilham/gold-store-demo/pkg/models"
)

func SaveTopupGold(cfg *configs.Config, data *models.BaseRequest) error {
	p, err := GetPrice(cfg)
	if err != nil {
		return err
	}

	if err := validateTopup(p.TopupPrice, data.Price, data.Gold); err != nil {
		return err
	}

	if err := SaveToAccount(cfg, &models.AccountData{
		AccountNumber: data.AccountNumber,
		Balance:       data.Gold,
	}, "topup"); err != nil {
		return err
	}

	if err := SaveToTransaction(cfg, &models.TransactionData{
		Gold: data.Gold,
		Type: "topup",
	}, data.AccountNumber); err != nil {
		return err
	}

	return nil
}

func validateTopup(topupPrice, price, gold float64) error {
	if price != topupPrice {
		return errors.New("harga doesn't match with current topup price")
	}

	n := gold / 0.001
	if n != float64(int(n)) || gold == 0 {
		return errors.New("minimum top-up is multiply of 0.001")
	}

	return nil
}
