package services

import (
	"encoding/json"
	"linkjo/app/models"
	"linkjo/config"
	"log"
	"os"
	"path/filepath"
)

// GetAllProducts mengambil semua produk dengan optional filter tenant_id atau category_id
func GetAllProducts(tenantID *uint, categoryID *uint) ([]models.Product, error) {
	var products []models.Product
	query := config.DB // Ambil koneksi database

	// Filter berdasarkan tenant_id jika ada
	if tenantID != nil {
		query = query.Where("tenant_id = ?", *tenantID)
	}

	// Filter berdasarkan category_id jika ada
	if categoryID != nil {
		query = query.Where("category_id = ?", *categoryID)
	}

	// Fetch data dengan relasi kategori
	err := query.Preload("Category").Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
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

func GetCategories() ([]models.Category, error) {
	basePath, err := os.Getwd()
	if err != nil {
		log.Println("Error getting current directory:", err)
		return nil, err
	}

	jsonPath := filepath.Join(basePath, "asset", "categories.json")

	data, err := os.ReadFile(jsonPath)
	if err != nil {
		log.Println("Error reading JSON file:", err)
		return nil, err
	}

	var categories []models.Category
	err = json.Unmarshal(data, &categories)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return nil, err
	}

	return categories, nil
}
