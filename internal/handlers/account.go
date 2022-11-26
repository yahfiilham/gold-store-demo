package handlers

import (
	"errors"

	"github.com/yahfiilham/gold-store-demo/configs"
	"github.com/yahfiilham/gold-store-demo/pkg/models"
	"gorm.io/gorm"
)

func SaveToAccount(cfg *configs.Config, data *models.AccountData, saveMode string) error {
	ac, err := GetAccount(cfg, data.AccountNumber)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		switch saveMode {
		case "topup":
			ac.Balance += data.Balance
		case "buyback":
			ac.Balance -= data.Balance
		}
		return cfg.DB.Model(ac).Where("id = ?", ac.ID).Updates(&ac).Error
	}

	m := &models.Account{
		AccountData: *data,
	}

	return cfg.DB.Model(m).Save(&m).Error
}

func GetAccount(cfg *configs.Config, accountNo string) (ac *models.Account, err error) {
	return ac, cfg.DB.Model(ac).Where("account_number = ?", accountNo).First(&ac).Error
}
