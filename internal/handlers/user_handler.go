package handlers

import (
	"context"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgtype"
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

func Register(c fiber.Ctx) error {
	var userReq entities.UserRegisterRequest
	var userRes entities.UserRegisterResponse

	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)

	defer cancel()

	if err := c.Bind().Body(&userReq); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(pkg.Response{
			Status:  500,
			Message: fiber.ErrInternalServerError.Message,
			Data:    err.Error(),
		})
	}

	if errors := userReq.Validate(); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.Response{
			Status:  409,
			Message: fiber.ErrBadRequest.Message,
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
		Id:        row.ID,
		Name:      row.Name,
		Email:     row.Email,
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
	}

	return c.Status(fiber.StatusOK).JSON(pkg.Response{
		Status:  201,
		Data:    userRes,
		Message: "successfully create data user",
	})
}

func Login(c fiber.Ctx) error {
	var userReq entities.UserLoginRequest
	var userRes entities.UserLoginResponse

	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	if err := c.Bind().Body(&userReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.Response{
			Status:  400,
			Message: fiber.ErrBadRequest.Message,
		})
	}

	if errors := userReq.Validate(); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.Response{
			Status:  400,
			Message: fiber.ErrBadRequest.Message,
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
			Message: fiber.ErrNotFound.Message,
		})
	}

	checked := CheckPasswordHash(pass, result.Password)
	if !checked {
		return c.Status(fiber.StatusUnauthorized).JSON(pkg.Response{
			Status:  401,
			Message: fiber.ErrUnauthorized.Message,
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
			Message: fiber.ErrInternalServerError.Message,
		})
	}

	userRes = entities.UserLoginResponse{
		Id:        result.ID,
		Name:      result.Name,
		Email:     result.Email,
		Token:     tokenString,
		CreatedAt: result.CreatedAt.Time,
		UpdatedAt: result.UpdatedAt.Time,
	}

	return c.Status(fiber.StatusOK).JSON(pkg.Response{
		Status:  200,
		Message: "login successful",
		Data:    userRes,
	})
}

func FindAllUser(c fiber.Ctx) error {
	db, _ := conf.NewPool()
	defer db.Close()

	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)

	defer cancel()

	q := users.New(db)

	result, err := q.FindAllUser(ctx)

	if err != nil {
		return c.Status(500).JSON(pkg.Response{
			Status:  500,
			Message: fiber.ErrInternalServerError.Error(),
		})
	}

	userRes := make([]entities.User, len(result))
	for i, row := range result {
		userRes[i] = entities.User{
			Id:        int(row.ID),
			Name:      row.Name,
			Email:     row.Email,
			CreatedAt: row.CreatedAt.Time,
			UpdatedAt: row.UpdatedAt.Time,
		}
	}

	if len(userRes) == 0 {
		userRes = []entities.User{}
	}

	return c.Status(fiber.StatusOK).JSON(pkg.Response{
		Status:  200,
		Message: "Successfuly find all users",
		Data:    userRes,
	})

}

func FindUserById(c fiber.Ctx) error {
	var userRes entities.User

	db, _ := conf.NewPool()
	defer db.Close()

	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)

	defer cancel()

	param := c.Params("id")
	id, _ := strconv.Atoi(param)

	q := users.New(db)

	result, err := q.FindUserById(ctx, int32(id))

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(pkg.Response{
			Status:  404,
			Message: fiber.ErrNotFound.Message,
		})
	}

	userRes = entities.User{
		Id:        int(result.ID),
		Name:      result.Name,
		Email:     result.Email,
		CreatedAt: userRes.CreatedAt,
		UpdatedAt: userRes.UpdatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(pkg.Response{
		Status:  200,
		Message: "successfully get data",
		Data:    userRes,
	})
}

func UpdateUser(c fiber.Ctx) error {
	var userReq entities.UpdateUserRequest
	var userRes entities.UpdateUserResponse

	db, _ := conf.NewPool()
	defer db.Close()
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	if err := c.Bind().Body(&userReq); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(pkg.Response{
			Status:  500,
			Message: fiber.ErrInternalServerError.Message,
		})
	}

	if errors := userReq.Validate(); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.Response{
			Status:  409,
			Message: fiber.ErrBadRequest.Message,
			Data:    errors,
		})
	}

	param := c.Params("id")
	id, _ := strconv.Atoi(param)

	q := users.New(db)

	updatedAt := time.Now()
	err := q.UpdateUserById(ctx, users.UpdateUserByIdParams{
		ID:        int32(id),
		Name:      userReq.Name,
		Email:     userReq.Email,
		UpdatedAt: pgtype.Timestamptz{Time: updatedAt, Valid: true},
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(pkg.Response{
			Status:  500,
			Message: fiber.ErrInternalServerError.Message,
		})
	}

	result, _ := q.FindUserById(ctx, int32(id))

	userRes = entities.UpdateUserResponse{
		Id:        result.ID,
		Name:      result.Name,
		Email:     result.Email,
		CreatedAt: result.CreatedAt.Time,
		UpdatedAt: result.UpdatedAt.Time,
	}

	return c.Status(fiber.StatusOK).JSON(pkg.Response{
		Status:  fiber.StatusOK,
		Message: "Successfully update data user",
		Data:    userRes,
	})
}

func DeleteUser(c fiber.Ctx) error {
	var userRes entities.User

	param := c.Params("id")
	id, _ := strconv.Atoi(param)

	db, _ := conf.NewPool()
	defer db.Close()

	q := users.New(db)
	row, err := q.FindUserById(c.Context(), int32(id))
	userRes = entities.User{
		Id:        int(row.ID),
		Name:      row.Name,
		Email:     row.Email,
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
	}

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(pkg.Response{
			Status:  404,
			Message: fiber.ErrNotFound.Message,
		})
	}

	err = q.DeleteUserById(c.Context(), int32(id))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(pkg.Response{
			Status:  200,
			Message: fiber.ErrInternalServerError.Message,
		})
	}

	return c.Status(fiber.StatusOK).JSON(pkg.Response{
		Status:  200,
		Message: "Successfully delete user",
		Data:    userRes,
	})

}
