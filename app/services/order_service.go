package services

import (
	"linkjo/app/models"

	"linkjo/config"
)

func CreateOrder(order *models.Order) (*models.Order, error) {
	err := config.DB.Create(order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}
