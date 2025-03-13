package controllers

import (
	"fmt"
	"linkjo/app/models"
	"linkjo/app/services"
	"linkjo/app/validators"

	"strconv"

	"github.com/gofiber/fiber/v2"
)

type StatisticsResponse struct {
	TotalOrders        uint    `json:"total_orders"`
	TotalRevenue       float64 `json:"total_revenue"`
	AverageOrder       float64 `json:"average_order"`
	ProductMostOrdered string  `json:"product_most_ordered"`
}

func GetProducts(c *fiber.Ctx) error {
	var tenantID *uint

	tID, ok := c.Locals("tenant_id").(string)
	if ok {
		id, err := strconv.ParseUint(tID, 10, 32)
		if err == nil {
			t := uint(id)
			tenantID = &t
		}
	}

	products, err := services.GetAllProducts(tenantID)
	if err != nil {
		response := models.APIResponse{
			Success: false,
			Message: "Failed to get products",
			Data:    nil,
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	message := "Products retrieved successfully"
	if len(products) == 0 {
		message = "No products found"
	}

	response := models.APIResponse{
		Success: true,
		Message: message,
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

	categoryIDStr := c.FormValue("category_id")

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

	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.APIResponse{
			Success: false,
			Message: "Invalid categoryID format",
			Data:    nil,
		})
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

	imagePaths := c.Locals("image_paths").(map[string][]string)
	fmt.Println("Image Paths:", imagePaths)
	// Image Paths: map[image:[./uploads/1741009398-1000041047.jpg]]

	product := models.Product{
		TenantID:    uint(tid),
		CategoryID:  uint(categoryID),
		Name:        req.Name,
		Price:       req.Price,
		Stock:       req.Stock,
		Unit:        req.Unit,
		Image:       imagePaths["image"][0],
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

func GetStatistics(c *fiber.Ctx) error {
	var tenant *uint

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

	tenantUidInt := uint(tid)
	tenant = &tenantUidInt

	totalOrders, err := services.GetTotalOrdersByUserID(tenant)
	if err != nil {
		response := models.APIResponse{
			Success: false,
			Message: "Failed to get total orders",
			Data:    nil,
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	totalRevenue, err := services.GetTotalRevenueByUserID(tenant)
	if err != nil {
		response := models.APIResponse{
			Success: false,
			Message: "Failed to get total revenue",
			Data:    nil,
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	averageOrder, err := services.GetAverageOrderAmountByUserID(tenant)
	if err != nil {
		response := models.APIResponse{
			Success: false,
			Message: "Failed to get average order amount",
			Data:    nil,
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	productMostOrdered, err := services.GetProductMostOrderedByUserID(tenant)
	if err != nil {
		response := models.APIResponse{
			Success: false,
			Message: "Failed to get product most ordered",
			Data:    nil,
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	bodyResponse := StatisticsResponse{
		TotalOrders:        uint(totalOrders),
		TotalRevenue:       totalRevenue,
		AverageOrder:       averageOrder,
		ProductMostOrdered: productMostOrdered,
	}
	response := models.APIResponse{
		Success: true,
		Message: "Statistics retrieved successfully",
		Data:    bodyResponse,
	}
	return c.JSON(response)
}
