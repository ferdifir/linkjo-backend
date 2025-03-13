package controllers

import (
	"linkjo/app/models"
	"linkjo/app/services"
	"linkjo/app/validators"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateOrder(c *fiber.Ctx) error {
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

	tID, ok := c.Locals("tenant_id").(string)
	if ok {
		id, err := strconv.ParseUint(tID, 10, 32)
		if err == nil {
			t := uint(id)
			tenantID = &t
		}
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

func GetOrders(c *fiber.Ctx) error {
	var tenantID *uint

	tID, ok := c.Locals("tenant_id").(string)
	if ok {
		id, err := strconv.ParseUint(tID, 10, 32)
		if err == nil {
			t := uint(id)
			tenantID = &t
		}
	}

	orders, err := services.GetOrdersByUserID(tenantID)
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
		Message: "Orders retrieved successfully",
		Data:    orders,
	}
	return c.JSON(response)
}

func UpdatePaymentStatus(c *fiber.Ctx) error {
	var req validators.OrderRequest

	if err := c.BodyParser(&req); err != nil {
		response := models.APIResponse{
			Success: false,
			Message: "Invalid request",
			Data:    nil,
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		response := models.APIResponse{
			Success: false,
			Message: "Invalid order ID",
			Data:    nil,
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if err := services.UpdatePaymentStatus(uint(id), req.PaymentStatus); err != nil {
		response := models.APIResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := models.APIResponse{
		Success: true,
		Message: "Payment status updated successfully",
		Data:    nil,
	}
	return c.JSON(response)
}
