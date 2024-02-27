package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type HealthcheckHandler struct {
	Validate *validator.Validate
}

func NewHealthcheckHandler() *HealthcheckHandler {
	return &HealthcheckHandler{}
}

func (h *HealthcheckHandler) ServeHTTP(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}
