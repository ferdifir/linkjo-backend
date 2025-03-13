package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"linkjo/app/models"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nfnt/resize"
)

func UploadImagesMiddleware(c *fiber.Ctx) error {
	savedNames := make(map[string][]string)

	form, err := c.MultipartForm()
	if err != nil {
		return c.Next()
	}

	for fieldName, files := range form.File {
		var fieldSavedNames []string

		for _, file := range files {
			allowedExtensions := map[string]bool{
				"image/jpeg": true,
				"image/jpg":  true,
				"image/png":  true,
			}
			if !allowedExtensions[file.Header.Get("Content-Type")] {
				return c.Status(fiber.StatusBadRequest).JSON(models.APIResponse{
					Success: false,
					Message: fmt.Sprintf("Invalid format for field %s. Only JPEG, JPG, and PNG files are allowed.", fieldName),
					Data:    nil,
				})
			}

			ext := filepath.Ext(file.Filename)
			if ext == "" {
				ext = ".jpg"
			}

			savedName := fmt.Sprintf("%d%s", time.Now().UnixNano(), strings.ToLower(ext))

			if file.Size > 2*1024*1024 {
				compressedFile, err := compressImage(file, ext)
				if err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(models.APIResponse{
						Success: false,
						Message: "Failed to compress image.",
						Data:    nil,
					})
				}
				file = compressedFile
			}

			uploadURL := "https://files.linkjo.my.id/upload"
			serverFileName, err := uploadToExternalAPI(uploadURL, file, savedName)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(models.APIResponse{
					Success: false,
					Message: "Failed to upload image to external API.",
					Data:    nil,
				})
			}

			fieldSavedNames = append(fieldSavedNames, serverFileName)
		}

		savedNames[fieldName] = fieldSavedNames
	}
	fmt.Println("savedNames: ", savedNames)
	c.Locals("image_paths", savedNames)

	return c.Next()
}

func compressImage(file *multipart.FileHeader, ext string) (*multipart.FileHeader, error) {
	srcFile, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer srcFile.Close()

	var img image.Image
	if ext == ".png" {
		img, err = png.Decode(srcFile)
	} else {
		img, err = jpeg.Decode(srcFile)
	}
	if err != nil {
		return nil, err
	}

	newImg := resize.Resize(1024, 0, img, resize.Lanczos3)

	var buf bytes.Buffer
	if ext == ".png" {
		err = png.Encode(&buf, newImg)
	} else {
		err = jpeg.Encode(&buf, newImg, &jpeg.Options{Quality: 80})
	}
	if err != nil {
		return nil, err
	}

	tmpFile, err := os.CreateTemp("", "compressed-*"+ext)
	if err != nil {
		return nil, err
	}
	defer tmpFile.Close()

	_, err = io.Copy(tmpFile, &buf)
	if err != nil {
		return nil, err
	}

	compressedFile := &multipart.FileHeader{
		Filename: file.Filename,
		Size:     int64(buf.Len()),
		Header:   file.Header,
	}
	return compressedFile, nil
}

func uploadToExternalAPI(apiURL string, file *multipart.FileHeader, savedName string) (string, error) {
	fileContent, err := file.Open()
	if err != nil {
		log.Println("Failed to open file:", err)
		return "", err
	}
	defer fileContent.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("files", savedName)
	if err != nil {
		log.Println("Failed to create form file:", err)
		return "", err
	}

	_, err = io.Copy(part, fileContent)
	if err != nil {
		log.Println("Failed to copy file content:", err)
		return "", err
	}
	writer.Close()

	req, err := http.NewRequest("POST", apiURL, body)
	if err != nil {
		log.Println("Failed to create request:", err)
		return "", err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Failed to send request:", err)
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Files []struct {
			SavedName string `json:"saved_name"`
		} `json:"files"`
		Message string `json:"message"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println("Failed to decode response:", err)
		return "", err
	}

	if len(result.Files) == 0 {
		log.Println("No files returned from API")
		return "", fmt.Errorf("no files returned from API")
	}

	return result.Files[0].SavedName, nil
}
