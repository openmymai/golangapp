package main

import (
	"log"
	"runtime/pprof"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Static("/", "./nextjs/dist")

	app.Get("/api", func(c *fiber.Ctx) error {
		profile := pprof.Lookup("allocs")
		return profile.WriteTo(c, 1)
	})

	log.Fatal(app.Listen(":8080"))
}
