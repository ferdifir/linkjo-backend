package middlewares

import (
	"fmt"
	"linkjo/app/models"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func UploadImagesMiddleware(c *fiber.Ctx) error {
	imagePaths := make(map[string][]string)

	form, err := c.MultipartForm()
	if err != nil {
		return c.Next()
	}

	for fieldName, files := range form.File {
		var fieldImages []string

		for _, file := range files {
			allowedExtensions := map[string]bool{
				"image/jpeg": true,
				"image/jpg":  true,
				"image/png":  true,
			}
			if !allowedExtensions[file.Header.Get("Content-Type")] {
				response := models.APIResponse{
					Success: false,
					Message: fmt.Sprintf("Invalid format for field %s. Only JPE, JPG, and PNG files are allowed.", fieldName),
					Data:    nil,
				}
				return c.Status(fiber.StatusBadRequest).JSON(response)
			}

			const maxSize = 2 * 1024 * 1024
			if file.Size > maxSize {
				response := models.APIResponse{
					Success: false,
					Message: fmt.Sprintf("File %s exceeds 2MB limit.", file.Filename),
					Data:    nil,
				}
				return c.Status(fiber.StatusBadRequest).JSON(response)
			}

			filename := fmt.Sprintf("%d-%s", time.Now().Unix(), strings.ReplaceAll(file.Filename, " ", "_"))
			filePath := fmt.Sprintf("./uploads/%s", filename)

			if err := c.SaveFile(file, filePath); err != nil {
				response := models.APIResponse{
					Success: false,
					Message: "Failed to save image.",
					Data:    nil,
				}
				return c.Status(fiber.StatusInternalServerError).JSON(response)
			}

			fieldImages = append(fieldImages, filePath)
		}

		imagePaths[fieldName] = fieldImages
	}

	c.Locals("image_paths", imagePaths)

	return c.Next()
}
