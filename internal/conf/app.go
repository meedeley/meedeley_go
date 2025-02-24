// internal/conf/conf.go
package conf

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/joho/godotenv"
)

func ProvideApp() (*fiber.App, error) {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Warning: Could not load .env file: %v", err)
	}

	app := fiber.New(fiber.Config{
		AppName:     os.Getenv("APP_NAME"),
		IdleTimeout: time.Hour * 1,
	})

	if err := setupLogger(app); err != nil {
		return nil, err
	}

	setUpGracefulShutdown(app)

	return app, nil
}

func setupLogger(app *fiber.App) error {
	if err := os.MkdirAll("tmp", 0755); err != nil {
		return err
	}

	logFile, err := os.OpenFile("tmp/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	app.Use(logger.New(logger.Config{
		Format:     "[LOG] ${time} ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "2006/01/2 15:04:05",
		Output:     logFile,
	}))

	return nil
}

func setUpGracefulShutdown(app *fiber.App) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		fmt.Println("\n🚨 Gracefully shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		shutdownDone := make(chan struct{})

		go func() {
			if err := app.ShutdownWithContext(ctx); err != nil {
				log.Printf("Shutdown error: %v", err)
			}
			close(shutdownDone)
		}()

		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for i := 5; i > 0; i-- {
			select {
			case <-ticker.C:
				fmt.Printf("⏳ Tersisa %d detik...\n", i)
			case <-shutdownDone:
				fmt.Println("✅ Shutdown completed early")
				return
			}
		}
		select {
		case <-shutdownDone:
			fmt.Println("✅ Shutdown completed")
		case <-ctx.Done():
			fmt.Println("⏰ Shutdown timeout")
		}
	}()
}
