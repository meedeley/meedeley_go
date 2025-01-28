package conf

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func RunAppWithGracefulShutdown() *fiber.App {

	err := godotenv.Load(".env")

	if err != nil {
		log.Printf("Warning: Could not load .env file: %v", err)
	}

	app := fiber.New(fiber.Config{
		AppName:       os.Getenv("APP_NAME"),
		StrictRouting: true,
		ErrorHandler:  fiber.DefaultErrorHandler,
		IdleTimeout:   time.Hour * 1,
	})

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		fmt.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	if err := app.Listen(":3000"); err != nil {
		log.Panic(err)
	}

	fmt.Println("Running cleanup tasks...")

	return app
}
