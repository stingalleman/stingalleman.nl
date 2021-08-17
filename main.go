package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/joho/godotenv/autoload"
	"github.com/stingalleman/stingalleman.dev/bedtime"
)

func main() {
	rootApp := fiber.New()

	rootApp.Use(logger.New())

	rootApp.Static("/", "./static")

	if os.Getenv("PORT") == "" {
		panic("no .env")
	}
	err := rootApp.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		panic(err)
	}
}
