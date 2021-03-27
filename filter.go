package filter

import (
	"github.com/gofiber/fiber/v2"
)

func New(config ...Config) fiber.Handler {
	// set default config
	cfg := configDefault(config...)

	return func(c *fiber.Ctx) error {
		// check condition
		if cfg.ShouldFilter(c) {
			return cfg.DoFilter(c)
		}

		// skip filter
		return nil
	}
}
