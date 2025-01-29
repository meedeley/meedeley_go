package conf

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func RunApp() *fiber.App {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Warning: Could not load .env file: %v", err)
	}

	app := fiber.New(fiber.Config{
		AppName:     os.Getenv("APP_NAME"),
		IdleTimeout: time.Hour * 1,
	})

	app.Use(cors.New())

	setUpGracefulShutdown(app)

	return app
}

func setUpGracefulShutdown(app *fiber.App) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		fmt.Println("Gracefully shutting down...")

		time.Sleep(3 * time.Second)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := app.ShutdownWithContext(ctx); err != nil {
			log.Printf("Error during shutdown: %v", err)
		}

		select {
		case <-ctx.Done():
			if ctx.Err() == context.DeadlineExceeded {
				fmt.Println("Shutdown timed out after 5 seconds, forcing exit...")
			} else {
				fmt.Println("Clean shutdown completed")
			}
		}

		fmt.Println("Cleanup tasks completed")
		os.Exit(0)
	}()
}
