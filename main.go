package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"net/http"
)

type (
	User struct {
		Name string `json:"name" validate:"required,min=5,max=20"` // Required field, min 5 char long max 20
		Age  int    `json:"age" validate:"required,numeric"`       // Required field, and client needs to implement our 'teener' tag format which we'll see later
	}

	XValidator struct {
		validator *validator.Validate
	}
)

var validate = validator.New()

func (v XValidator) Validate(data interface{}) error {
	return validate.Struct(data)
}

func main() {
	myValidator := &XValidator{
		validator: validate,
	}

	app := fiber.New()
	app.Use(logger.New())

	app.Get(
		"/", func(c *fiber.Ctx) error {
			user := &User{
				Name: c.Query("name"),
				Age:  c.QueryInt("age"),
			}

			// Validation
			if err := myValidator.Validate(user); err != nil {
				return fiber.NewError(http.StatusBadRequest, err.Error())
			}

			// Logic, validated with success
			return c.Status(http.StatusCreated).JSON(user)
		},
	)

	log.Fatal(app.Listen(":8080"))
}
