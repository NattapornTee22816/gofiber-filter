package filter

import (
	"github.com/gofiber/fiber/v2"
	"regexp"
	"strings"
)

type Config struct {
	// url pattern for filter (use regex pattern)
	// default ["*"] all url
	// url pattern not should set to should set
	// like /api/**/* => /api/* (not use **)
	UrlPattern []string
	// case sensitive
	// when true append (?i) to prefix regex pattern
	// default false
	CaseSensitive bool
	// http method
	// default GET, PUT, POST, DELETE
	MethodPattern []string
	// condition for filter
	// when pattern (url and method) match (return true) will execute DoFilter
	// default DefaultShouldFilter
	ShouldFilter func(c *fiber.Ctx) bool
	// actions for filter
	DoFilter func(c *fiber.Ctx) error
}

var defaultUrlPattern = []string{"*"}
var defaultMethodPattern = []string{
	fiber.MethodGet,
	fiber.MethodPut,
	fiber.MethodPost,
	fiber.MethodDelete,
}

func configDefault(config ...Config) Config {
	var cfg Config
	if len(config) > 0 {
		cfg = config[0]
	} else {
		cfg = Config{}
	}

	// default url pattern
	if cfg.UrlPattern == nil || len(cfg.UrlPattern) == 0 {
		cfg.UrlPattern = defaultUrlPattern
	}

	// default method pattern
	if cfg.MethodPattern == nil || len(cfg.MethodPattern) == 0 {
		cfg.MethodPattern = defaultMethodPattern
	}

	// default shouldFilter
	if cfg.ShouldFilter == nil {
		cfg.ShouldFilter = cfg.DefaultShouldFilter
	}

	// default DoFilter
	if cfg.DoFilter == nil {
		cfg.DoFilter = defaultDoFilter
	}

	return cfg
}

func (cfg *Config) DefaultShouldFilter(c *fiber.Ctx) bool {
	if matched := cfg.matchMethod(c.Method()); !matched {
		return false
	}

	path := string(c.Request().URI().Path())
	if matched := match(cfg.UrlPattern, path, cfg.CaseSensitive); !matched {
		return false
	}

	return true
}

func defaultDoFilter(c *fiber.Ctx) error {
	return nil
}

func (cfg *Config) matchMethod(method string) bool {
	// check * in list
	// when found will ignore others
	if match(cfg.MethodPattern, "*", false) {
		return true
	}
	return match(cfg.MethodPattern, method, true)
}

func match(patterns []string, match string, fold bool) bool {
	for _, pattern := range patterns {
		if fold && strings.EqualFold(pattern, match) {
			return true
		} else if pattern == match {
			return true
		}

		if matched := matchStep(pattern, match, fold); matched {
			return true
		}
	}
	return false
}

func matchStep(p string, m string, fold bool) bool {
	// check with regex pattern
	pRegexPattern := strings.ReplaceAll(p, "*", ".*")
	if fold {
		pRegexPattern = "(?i)" + pRegexPattern
	}

	pRegex := regexp.MustCompile(pRegexPattern)
	if matched := pRegex.MatchString(m); matched {
		return true
	}

	return false
}
