package services

import (
	"linkjo/app/models"
	"linkjo/config"
)

// get total orders based user id
func GetTotalOrdersByUserID(userID *uint) (int64, error) {
	var totalOrders int64
	result := config.DB.Model(&models.Order{}).Where("user_id = ?", userID).Count(&totalOrders)
	if result.Error != nil {
		return 0, result.Error
	}
	return totalOrders, nil
}

// get total revenue based user id
func GetTotalRevenueByUserID(userID *uint) (float64, error) {
	var totalRevenue float64
	result := config.DB.Model(&models.Order{}).Where("user_id = ?", userID).Select("sum(total_price)").Pluck("sum(total_price)", &totalRevenue)
	if result.Error != nil {
		return 0, result.Error
	}
	return totalRevenue, nil
}

// average order amount based user id
func GetAverageOrderAmountByUserID(userID *uint) (float64, error) {
	var totalOrders int64
	var totalRevenue float64
	result := config.DB.Model(&models.Order{}).Where("user_id = ?", userID).Count(&totalOrders)
	if result.Error != nil {
		return 0, result.Error
	}

	result = config.DB.Model(&models.Order{}).Where("user_id = ?", userID).Select("sum(total_price)").Pluck("sum(total_price)", &totalRevenue)
	if result.Error != nil {
		return 0, result.Error
	}

	if totalOrders > 0 {
		return totalRevenue / float64(totalOrders), nil
	} else {
		return 0, nil
	}
}

// get product the most ordered based user id
func GetProductMostOrderedByUserID(userID *uint) (string, error) {
	var productMostOrdered uint

	result := config.DB.Table("order_details").
		Select("order_details.product_id").
		Joins("JOIN orders ON order_details.order_id = orders.id").
		Where("orders.user_id = ? AND orders.deleted_at IS NULL", userID).
		Group("order_details.product_id").
		Order("COUNT(order_details.product_id) DESC").
		Limit(1).
		Pluck("order_details.product_id", &productMostOrdered)

	if result.Error != nil {
		return "", result.Error
	}

	// get product name
	var productName string
	result = config.DB.Model(&models.Product{}).Where("id = ?", productMostOrdered).Pluck("name", &productName)
	if result.Error != nil {
		return "", result.Error
	}

	return productName, nil
}
