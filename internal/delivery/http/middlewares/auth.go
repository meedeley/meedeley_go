// internal/middleware/auth.go
package middlewares

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/meedeley/go-launch-starter-code/internal/conf"
	"github.com/meedeley/go-launch-starter-code/pkg"
)

type AuthMiddleware struct {
	jwtSecret string
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{
		jwtSecret: conf.JwtSecret(),
	}
}

func (a *AuthMiddleware) GuestOnly() fiber.Handler {
	return func(c fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Next()
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
			}
			return []byte(a.jwtSecret), nil
		})

		if err == nil && token.Valid {
			return c.Status(fiber.StatusForbidden).JSON(pkg.Response{
				Status:  403,
				Message: "Access denied for logged-in users",
			})
		}

		return c.Next()
	}
}

func (a *AuthMiddleware) Protected() fiber.Handler {
	return func(c fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(pkg.Response{
				Status:  401,
				Message: "Missing or invalid Authorization header",
			})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
			}
			return []byte(a.jwtSecret), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(pkg.Response{
				Status:  401,
				Message: "Invalid or expired token",
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(pkg.Response{
				Status:  401,
				Message: "Invalid token claims",
			})
		}

		if exp, ok := claims["exp"].(float64); ok {
			if int64(exp) < time.Now().Unix() {
				return c.Status(fiber.StatusUnauthorized).JSON(pkg.Response{
					Status:  401,
					Message: "Token Expired",
				})
			}
		}

		c.Locals("user", claims)
		return c.Next()
	}
}
