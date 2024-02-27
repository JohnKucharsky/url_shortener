package handlers

import (
	"github.com/JohnKucharsky/url_shortener/store"
	"github.com/gofiber/fiber/v2"
	"math/rand"
	"net/http"
)

type CreateShortURLHandler struct {
	shortURLStore store.ShortURLStore
}

type CreateShortURLHandlerParams struct {
	ShortURLStore store.ShortURLStore
}

func NewCreateShortURLHandler(params CreateShortURLHandlerParams) *CreateShortURLHandler {
	return &CreateShortURLHandler{shortURLStore: params.ShortURLStore}
}

func generateSlug() string {
	const charSet = "abcdefghiklmnoprstuvwxyz0123456789"
	result := make([]byte, 6)

	for i := range result {
		result[i] = charSet[rand.Intn(len(charSet))]
	}

	return string(result)
}

func (h *CreateShortURLHandler) ServeHTTP(c *fiber.Ctx) error {
	var requestData = struct {
		Destination string `json:"destination"`
	}{}

	if err := c.BodyParser(&requestData); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	slug := generateSlug()

	createdShortURL, _ := h.shortURLStore.CreateShortURL(
		store.CreateShortURLParams{
			Destination: requestData.Destination,
			Slug:        slug,
		},
	)
	return c.Status(http.StatusCreated).JSON(createdShortURL)
}
