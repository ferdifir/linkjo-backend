package controllers

import (
	"linkjo/app/models"
	"linkjo/app/services"
	"linkjo/app/validators"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	var tenantID *uint
	var categoryID *uint

	if tID := c.Query("tenant_id"); tID != "" {
		id, err := strconv.ParseUint(tID, 10, 32)
		if err == nil {
			t := uint(id)
			tenantID = &t
		}
	}

	if cID := c.Query("category_id"); cID != "" {
		id, err := strconv.ParseUint(cID, 10, 32)
		if err == nil {
			cat := uint(id)
			categoryID = &cat
		}
	}

	products, err := services.GetAllProducts(tenantID, categoryID)
	if err != nil {
		response := models.APIResponse{
			Success: false,
			Message: "Failed to get products",
			Data:    nil,
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := models.APIResponse{
		Success: true,
		Message: "Products retrieved successfully",
		Data:    products,
	}
	return c.JSON(response)
}

func GetProductByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		response := models.APIResponse{
			Success: false,
			Message: "Invalid product ID",
			Data:    nil,
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	product, err := services.GetProductByID(uint(id))
	if err != nil {
		response := models.APIResponse{
			Success: false,
			Message: "Product not found",
			Data:    nil,
		}
		return c.Status(fiber.StatusNotFound).JSON(response)
	}

	response := models.APIResponse{
		Success: true,
		Message: "Product retrieved successfully",
		Data:    product,
	}
	return c.JSON(response)
}

func CreateProduct(c *fiber.Ctx) error {
	var req validators.ProductRequest

	if err := c.BodyParser(&req); err != nil {
		response := models.APIResponse{
			Success: false,
			Message: "Invalid request",
			Data:    nil,
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	tenantID, ok := c.Locals("tenant_id").(string)
	if !ok || tenantID == "" {
		response := models.APIResponse{
			Success: false,
			Message: "Tenant ID tidak ditemukan",
			Data:    nil,
		}
		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}

	tid, err := strconv.ParseUint(tenantID, 10, 32)
	if err != nil {
		response := models.APIResponse{
			Success: false,
			Message: "Invalid tenantID format",
			Data:    nil,
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	product := models.Product{
		TenantID:    uint(tid),
		CategoryID:  req.CategoryID,
		Name:        req.Name,
		Price:       req.Price,
		Stock:       req.Stock,
		Unit:        req.Unit,
		Image:       req.Image,
		Description: req.Description,
	}

	createdProduct, err := services.CreateProduct(&product)
	if err != nil {
		response := models.APIResponse{
			Success: false,
			Message: "Failed to create product",
			Data:    nil,
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	response := models.APIResponse{
		Success: true,
		Message: "Product created successfully",
		Data:    createdProduct,
	}
	return c.JSON(response)
}

func GetCategories(c *fiber.Ctx) error {
	categories, err := services.GetCategories()
	if err != nil {
		response := models.APIResponse{
			Success: false,
			Message: "Failed to get categories",
			Data:    nil,
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := models.APIResponse{
		Success: true,
		Message: "Categories retrieved successfully",
		Data:    categories,
	}
	return c.JSON(response)
}
