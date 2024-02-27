package main

import (
	"fmt"
	"github.com/JohnKucharsky/url_shortener/handlers"
	"github.com/JohnKucharsky/url_shortener/store/dbstore"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"log/slog"
	"os"
	"os/signal"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		fmt.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	lo := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	app.Get(
		"/health", handlers.NewHealthcheckHandler().ServeHTTP,
	)

	shortURLStore := dbstore.NewShortURLStore(lo)

	app.Post(
		"/short_url",
		handlers.NewCreateShortURLHandler(
			handlers.CreateShortURLHandlerParams{
				ShortURLStore: shortURLStore,
			},
		).ServeHTTP,
	)

	if err := app.Listen(":8080"); err != nil {
		log.Panic(err)
	}

	fmt.Println("Running cleanup tasks...")
}
