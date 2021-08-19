package bedtime

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Init() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		tz := c.Query("tz")
		if tz == "" {
			tz = "Europe/Amsterdam"
		}

		bedtime := c.Query("time")
		if bedtime == "" {
			return fmt.Errorf("must suply time query (in 24h)\nexample: https://stingalleman.dev/bedtime?time=10:00")
		}

		times, err := CalculateBedTime(bedtime, tz)
		if err != nil {
			return err
		}

		c.JSON(times)
		return err
	})

	app.Get("/now", func(c *fiber.Ctx) error {
		tz := c.Query("tz")
		if tz == "" {
			tz = "Europe/Amsterdam"
		}

		times, err := CalculateBedtimeNow(tz)
		if err != nil {
			return err
		}

		c.JSON(times)
		return err
	})

	return app
}
