package handlers

import (
	"context"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/meedeley/go-launch-starter-code/db/models/users"
	"github.com/meedeley/go-launch-starter-code/internal/conf"
	"github.com/meedeley/go-launch-starter-code/internal/entities"
	"github.com/meedeley/go-launch-starter-code/internal/usecase"
	"github.com/meedeley/go-launch-starter-code/pkg"
)

type UserHandler struct {
	UseCase usecase.UserUseCase
}

func NewUserHandler(usecase usecase.UserUseCase) *UserHandler {
	return &UserHandler{UseCase: usecase}
}

func (uc *UserHandler) Register(c fiber.Ctx) error {
	var userReq entities.UserRegisterRequest

	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)

	defer cancel()

	if err := c.Bind().Body(&userReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.Response{
			Status:  400,
			Message: fiber.ErrInternalServerError.Message,
			Data:    err.Error(),
		})
	}

	if errors := userReq.Validate(); errors != nil {
		return pkg.NewErrorValidation(c, errors)
	}

	userRes, err := uc.UseCase.Register(ctx, userReq)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(pkg.Response{
			Status:  500,
			Message: fiber.ErrInternalServerError.Message,
			Data:    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(pkg.Response{
		Status:  201,
		Data:    userRes,
		Message: "successfully create data user",
	})
}

func (uc *UserHandler) Login(c fiber.Ctx) error {
	var userReq entities.UserLoginRequest

	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	if err := c.Bind().Body(&userReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.Response{
			Status:  400,
			Message: fiber.ErrBadRequest.Message,
			Data:    err.Error(),
		})
	}

	if errors := userReq.Validate(); errors != nil {
		pkg.NewErrorValidation(c, errors)
	}

	userRes, err := uc.UseCase.Login(ctx, userReq)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(pkg.Response{
			Status:  500,
			Message: fiber.StatusInternalServerError,
			Data:    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(pkg.Response{
		Status:  200,
		Message: "login successful",
		Data:    userRes,
	})
}

func (h *UserHandler) FindAll(c fiber.Ctx) error {
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
		var updatedAt *time.Time
		if row.UpdatedAt.Valid {
			updatedAt = &row.UpdatedAt.Time
		}
		userRes[i] = entities.User{
			Id:        int(row.ID),
			Name:      row.Name,
			Email:     row.Email,
			CreatedAt: row.CreatedAt.Time,
			UpdatedAt: updatedAt,
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

func (h *UserHandler) FindById(c fiber.Ctx) error {
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

	var updatedAt *time.Time
	if result.UpdatedAt.Valid {
		updatedAt = &result.UpdatedAt.Time
	}

	userRes = entities.User{
		Id:        int(result.ID),
		Name:      result.Name,
		Email:     result.Email,
		CreatedAt: result.CreatedAt.Time,
		UpdatedAt: updatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(pkg.Response{
		Status:  200,
		Message: "successfully get data",
		Data:    userRes,
	})
}

func Update(c fiber.Ctx) error {
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

func Delete(c fiber.Ctx) error {
	var userRes entities.User

	param := c.Params("id")
	id, _ := strconv.Atoi(param)

	db, _ := conf.NewPool()
	defer db.Close()

	q := users.New(db)
	row, err := q.FindUserById(c.Context(), int32(id))

	var updatedAt *time.Time
	if row.UpdatedAt.Valid {
		updatedAt = &row.UpdatedAt.Time
	}
	userRes = entities.User{
		Id:        int(row.ID),
		Name:      row.Name,
		Email:     row.Email,
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: updatedAt,
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
