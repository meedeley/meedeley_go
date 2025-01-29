package conf

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func JwtSecret(key string) string {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Error loading .enf file")
	}

	return os.Getenv(key)
}
