package handlers

import (
	"context"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/meedeley/go-launch-starter-code/db/models/users"
	"github.com/meedeley/go-launch-starter-code/internal/conf"
	"github.com/meedeley/go-launch-starter-code/internal/entity"
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
	var userReq entity.UserRegisterRequest

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
	var userReq entity.UserLoginRequest

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
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)

	defer cancel()

	userRes, err := h.UseCase.FindAll(ctx)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(pkg.Response{
			Status:  500,
			Message: fiber.ErrInternalServerError.Message,
			Data:    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(pkg.Response{
		Status:  200,
		Message: "Successfuly find all users",
		Data:    userRes,
	})

}

func (h *UserHandler) FindById(c fiber.Ctx) error {
	var userRes entity.User

	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)

	defer cancel()

	param := c.Params("id")
	id, _ := strconv.Atoi(param)

	userRes, err := h.UseCase.FindById(ctx, int32(id))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(pkg.Response{
			Status:  500,
			Message: fiber.ErrInternalServerError.Message,
			Data:    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(pkg.Response{
		Status:  200,
		Message: "successfully get data",
		Data:    userRes,
	})
}

func (h *UserHandler) Update(c fiber.Ctx) error {
	var userReq entity.UpdateUserRequest

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

	userRes, err := h.UseCase.Update(ctx, int32(id), userReq)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(pkg.Response{
			Status:  500,
			Message: fiber.ErrInternalServerError.Message,
			Data:    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(pkg.Response{
		Status:  fiber.StatusOK,
		Message: "Successfully update data user",
		Data:    userRes,
	})
}

func Delete(c fiber.Ctx) error {
	var userRes entity.User

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
	userRes = entity.User{
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
