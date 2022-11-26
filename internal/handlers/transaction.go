package handlers

import (
	"github.com/yahfiilham/gold-store-demo/configs"
	"github.com/yahfiilham/gold-store-demo/pkg/models"
)

func SaveToTransaction(cfg *configs.Config, data *models.TransactionData, accountNo string) error {
	ac, err := GetAccount(cfg, accountNo)
	if err != nil {
		return err
	}

	p, err := GetPrice(cfg)
	if err != nil {
		return err
	}

	m := &models.Transaction{
		TransactionData: *data,
		AccountData:     ac.AccountData,
		PriceData:       p.PriceData,
	}

	return cfg.DB.Model(m).Save(&m).Error
}

func GetMutation(cfg *configs.Config, accountNo string, startDate, endDate int64) (m []*models.Transaction, err error) {
	return m, cfg.DB.
		Model(&models.Transaction{}).
		Where("account_number = ? AND created_at >= ? AND updated_at <= ?", accountNo, startDate, endDate).
		Find(&m).Error
}
