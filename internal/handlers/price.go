package handlers

import (
	"github.com/yahfiilham/gold-store-demo/configs"
	"github.com/yahfiilham/gold-store-demo/pkg/models"
)

func SavePrice(cfg *configs.Config, data *models.PriceData) (*models.Price, error) {
	m := &models.Price{
		PriceData: *data,
	}

	return m, cfg.DB.Model(m).Save(&m).Error
}

func GetPrice(cfg *configs.Config) (p *models.Price, err error) {
	return p, cfg.DB.Model(p).Order("created_at DESC").First(&p).Error
}
