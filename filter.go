package filter

import (
	"github.com/gofiber/fiber/v2"
)

func New(config Config) fiber.Handler {
	// set default config
	config = configDefault(config)

	return func(c *fiber.Ctx) error {
		// check condition
		if config.ShouldFilter(c) {
			return config.DoFilter(c)
		}

		// skip filter
		return c.Next()
	}
}