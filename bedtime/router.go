package bedtime

import (
	"time"

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
			loc, err := time.LoadLocation(tz)
			if err != nil {
				return err
			}

			bedtime = time.Now().In(loc).Format("15:04")
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
