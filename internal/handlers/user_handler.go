package handlers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/meedeley/go-launch-starter-code/db/models/users"
	"github.com/meedeley/go-launch-starter-code/internal/conf"
	"github.com/meedeley/go-launch-starter-code/internal/entities"
	"github.com/meedeley/go-launch-starter-code/pkg"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Register(c *fiber.Ctx) error {
	var userReq entities.UserRegisterRequest
	var userRes entities.UserRegisterResponse

	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)

	defer cancel()

	if err := c.BodyParser(&userReq); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(pkg.Response{
			Status:  500,
			Message: "Failed to parse request body",
			Data:    err.Error(),
		})
	}

	if errors := userReq.Validate(); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.Response{
			Status:  409,
			Message: "failed to validate register request",
			Data:    errors,
		})
	}

	db, _ := conf.NewPool()
	defer db.Close()
	q := users.New(db)

	hashedPassword, _ := hashPassword(userReq.Password)

	row, err := q.InsertUser(ctx, users.InsertUserParams{
		Name:     userReq.Name,
		Email:    userReq.Email,
		Password: hashedPassword,
	})

	if err != nil {
		return err
	}

	userRes = entities.UserRegisterResponse{
		Id:    row.ID,
		Name:  row.Name,
		Email: row.Email,
	}

	return c.JSON(pkg.Response{
		Status:  201,
		Data:    userRes,
		Message: "successfully create data user",
	})
}

func Login(c *fiber.Ctx) error {
	var userReq entities.UserLoginRequest
	var userRes entities.UserLoginResponse

	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	if err := c.BodyParser(&userReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.Response{
			Status:  400,
			Message: "failed to parse request body",
		})
	}

	if errors := userReq.Validate(); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.Response{
			Status:  400,
			Message: "failed to validate login request",
			Data:    errors,
		})
	}

	db, _ := conf.NewPool()
	defer db.Close()
	q := users.New(db)

	email := userReq.Email
	pass := userReq.Password

	result, err := q.FindUserByEmail(ctx, email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(pkg.Response{
			Status:  404,
			Message: "user not found",
		})
	}

	checked := CheckPasswordHash(pass, result.Password)
	if !checked {
		return c.Status(fiber.StatusUnauthorized).JSON(pkg.Response{
			Status:  401,
			Message: "invalid password",
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	var claims jwt.MapClaims = token.Claims.(jwt.MapClaims)
	claims["user_id"] = result.ID
	claims["email"] = result.Email
	claims["name"] = result.Name
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte(conf.JwtSecret()))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(pkg.Response{
			Status:  500,
			Message: "failed to generate token",
		})
	}

	userRes = entities.UserLoginResponse{
		Id:    result.ID,
		Name:  result.Name,
		Email: result.Email,
		Token: tokenString,
	}

	return c.JSON(pkg.Response{
		Status:  200,
		Message: "login successful",
		Data:    userRes,
	})
}

func FindAllUser(c *fiber.Ctx) error {
	db, _ := conf.NewPool()
	defer db.Close()

	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)

	defer cancel()

	q := users.New(db)

	result, err := q.FindAllUser(ctx)

	if err != nil {
		return c.Status(500).JSON(pkg.Response{
			Status:  500,
			Message: "Internal servel error",
		})
	}

	if result == nil {
		result = []users.FindAllUserRow{}
	}

	return c.Status(200).JSON(pkg.Response{
		Status:  200,
		Message: "Successfuly find all users",
		Data:    result,
	})

}

func findUserById() {

}
