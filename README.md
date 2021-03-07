# gofiber-filter middleware

Filter middleware for [gofiber](https://github.com/gofiber/fiber)
inspire by [spring framework](https://spring.io/)

## Installation

```
go get -u github.com/NattapornTee22816/gofiber-filter
```

## Import

```
import (
    "github.com/gofiber/fiber/v2"
    filter "github.com/NattapornTee22816/gofiber-filter"
)
```

## Example

```
app := fiber.New()
app.Use(filter.New())
```

## Default Config

```
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
```
