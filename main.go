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

	bedtimeApp := bedtime.Init()
	rootApp.Mount("/bedtime", bedtimeApp)

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	err := rootApp.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}
}
