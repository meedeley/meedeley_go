package conf

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func JwtSecret() string {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Error loading .env file")
	}

	secret := os.Getenv("JWT_SECRET")

	return secret
}
