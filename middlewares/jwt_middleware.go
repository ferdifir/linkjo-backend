package middlewares

import (
	"fmt"
	"linkjo/app/models"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v4"
)

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			log.Println("DEBUG: Authorization header kosong")
			return c.Status(fiber.StatusUnauthorized).JSON(models.APIResponse{
				Success: false,
				Message: "Unauthorized: Token tidak ditemukan",
				Data:    nil,
			})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			log.Println("DEBUG: Format Authorization header salah")
			return c.Status(fiber.StatusUnauthorized).JSON(models.APIResponse{
				Success: false,
				Message: "Unauthorized: Format token salah",
				Data:    nil,
			})
		}

		tokenString := parts[1]
		secretKey := os.Getenv("JWT_SECRET")
		if secretKey == "" {
			secretKey = "linkjo"
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.Printf("DEBUG: Signing method tidak sesuai: %v", token.Header["alg"])
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			log.Printf("DEBUG: Token tidak valid: %v", err)
			return c.Status(fiber.StatusUnauthorized).JSON(models.APIResponse{
				Success: false,
				Message: "Unauthorized: Token tidak valid",
				Data:    nil,
			})
		}

		c.Locals("user", token)

		return c.Next()
	}
}

func ExtractTenantID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userToken := c.Locals("user")

		if userToken == nil {
			response := models.APIResponse{
				Success: false,
				Message: "Token tidak ditemukan",
				Data:    nil,
			}
			return c.Status(fiber.StatusUnauthorized).JSON(response)
		}

		token, ok := userToken.(*jwt.Token)
		if !ok {
			log.Printf("Error: Token format tidak valid. Tipe sebenarnya: %T", userToken)
			response := models.APIResponse{
				Success: false,
				Message: "Format token tidak valid",
				Data:    nil,
			}
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := models.APIResponse{
				Success: false,
				Message: "Klaim token tidak valid",
				Data:    nil,
			}
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		}

		tenantIDFloat, ok := claims["tenant_id"].(float64)
		if !ok {
			response := models.APIResponse{
				Success: false,
				Message: "Tenant ID tidak ditemukan di token",
				Data:    nil,
			}
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}

		tenantID := strconv.Itoa(int(tenantIDFloat))

		c.Locals("tenant_id", tenantID)

		return c.Next()
	}
}
