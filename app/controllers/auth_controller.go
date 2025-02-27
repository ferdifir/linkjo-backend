package controllers

import (
	"linkjo/app/models"
	"linkjo/app/services"
	"linkjo/app/validators"

	"github.com/gofiber/fiber/v2"
)

func RegisterUser(c *fiber.Ctx) error {
	var req validators.RegisterRequest

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

	user := models.User{
		OwnerName:  req.OwnerName,
		OutletName: req.OutletName,
		Email:      req.Email,
		Password:   req.Password,
		Phone:      req.Phone,
		Address:    req.Address,
		City:       req.City,
		Latitude:   req.Latitude,
		Longitude:  req.Longitude,
		IsActive:   req.IsActive,
	}

	createdUser, err := services.RegisterUser(user)
	if err != nil {
		response := models.APIResponse{
			Success: false,
			Message: "Failed to register user",
			Data:    nil,
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := models.APIResponse{
		Success: true,
		Message: "User registered successfully",
		Data:    createdUser,
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

func LoginUser(c *fiber.Ctx) error {
	var req validators.LoginRequest

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

	token, err := services.LoginUser(req.Email, req.Password)
	if err != nil {
		response := models.APIResponse{
			Success: false,
			Message: "Failed to login",
			Data:    nil,
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := models.APIResponse{
		Success: true,
		Message: "User login successfully",
		Data:    token,
	}

	return c.Status(fiber.StatusOK).JSON(response)

}
