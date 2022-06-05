package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/google/uuid"
)

func main() {
	// Fiber instance
	app := fiber.New()

	app.Use(cors.New())
	// Routes
	app.Static("/images", "./images")
	app.Post("/", handleFileupload)
	app.Delete("/:imageName", handleDeleteImage)

	// Start server
	log.Fatal(app.Listen(":8080"))
}

func handleFileupload(c *fiber.Ctx) error {
	file, err := c.FormFile("image")

	if err != nil {
		log.Println("Image upload error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "messge": "Server error", "data": nil})
	}
	uniqueId := uuid.New()
	filename := strings.Replace(uniqueId.String(), "-", "", -1)
	fileExt := strings.Split(file.Filename, ".")[1]
	image := fmt.Sprintf("%s.%s", filename, fileExt)
	err = c.SaveFile(file, fmt.Sprintf("./images/%s", image))

	if err != nil {
		log.Println("Image save error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}
	imageUrl := fmt.Sprintf("http://localhost:8080/images/%s", image)

	data := map[string]interface{}{
		"imageName": image,
		"imageUrl":  imageUrl,
		"header":    file.Header,
		"size":      file.Size,
	}

	return c.JSON(fiber.Map{"status": 201, "message": "Image upload successfully", "data": data})
}

func handleDeleteImage(c *fiber.Ctx) error {
	imageName := c.Params("imageName")
	err := os.Remove(fmt.Sprintf("./images/%s", imageName))
	if err != nil {
		log.Println(err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}
	return c.JSON(fiber.Map{"status": 201, "message": "Image deleted successfully", "data": nil})
}
