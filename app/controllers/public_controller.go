package controllers

import (
	"linkjo/app/models"
	"linkjo/app/services"
	"linkjo/app/validators"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetProductsByTenantID(c *fiber.Ctx) error {
	tenantID := c.Query("tenant_id")

	tID, err := strconv.ParseUint(tenantID, 10, 32)
	if err != nil {
		response := models.APIResponse{
			Success: false,
			Message: "Invalid tenantID format",
			Data:    nil,
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	tenantUidInt := uint(tID)
	tenant := &tenantUidInt

	products, err := services.GetAllProducts(tenant)
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

func CreateOrderByTenantID(c *fiber.Ctx) error {
	var tenantID *uint
	var req validators.OrderRequest

	if err := c.BodyParser(&req); err != nil {
		response := models.APIResponse{
			Success: false,
			Message: "Invalid request",
			Data:    nil,
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if err := validators.ValidateStruct(req); err != nil {
		response := models.APIResponse{
			Success: false,
			Message: "Invalid request",
			Data:    nil,
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	tID := c.Query("tenant_id")

	id, err := strconv.ParseUint(tID, 10, 32)
	if err == nil {
		t := uint(id)
		tenantID = &t
	}

	order, err := services.CreateOrder(req, tenantID)
	if err != nil {
		response := models.APIResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := models.APIResponse{
		Success: true,
		Message: "Order created successfully",
		Data:    order,
	}
	return c.JSON(response)
}
