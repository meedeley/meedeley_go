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

	app := fiber.New(fiber.Config{
		AppName:      os.Getenv("APP_NAME"),
		ErrorHandler: fiber.DefaultErrorHandler,
		IdleTimeout:  time.Hour * 1,
	})

	app.Use(cors.New())

	return setUpRouteWithGracefullShutdown(app)
}

func setUpRouteWithGracefullShutdown(app *fiber.App) *fiber.App {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Warning: Could not load .env file: %v", err)
	}

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

		<-ctx.Done()
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("Shutdown timed out after 5 seconds, forcing exit...")
		} else {
			fmt.Println("Clean shutdown completed")
		}

		fmt.Println("Cleanup tasks completed")
		os.Exit(0)
	}()

	return app
}
