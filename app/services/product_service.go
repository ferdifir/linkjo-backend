package services

import (
	"linkjo/app/models"
	"linkjo/config"
)

type ProductWithCategory struct {
	models.Product
	CategoryName string `json:"category_name"`
}

func GetAllProducts(tenantID *uint) (map[string][]ProductWithCategory, error) {
	var products []ProductWithCategory

	query := config.DB.
		Table("products").
		Select(`products.*, categories.name as category_name`).
		Joins("LEFT JOIN categories ON products.category_id = categories.id").
		Where("products.tenant_id = ?", tenantID)

	err := query.Find(&products).Error
	if err != nil {
		return nil, err
	}

	groupedProducts := make(map[string][]ProductWithCategory)

	for _, product := range products {
		category := product.CategoryName
		groupedProducts[category] = append(groupedProducts[category], product)
	}

	return groupedProducts, nil
}

// GetProductByID mengambil detail produk berdasarkan ID
func GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	err := config.DB.Preload("Category").First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func CreateProduct(product *models.Product) (*models.Product, error) {
	err := config.DB.Create(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func GetCategories() ([]models.Categories, error) {
	var categories []models.Categories

	err := config.DB.
		Preload("Children").
		Where("parent_id IS NULL").
		Find(&categories).Error

	if err != nil {
		return nil, err
	}
	return categories, nil
}
