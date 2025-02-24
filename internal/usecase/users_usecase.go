package usecase

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/meedeley/go-launch-starter-code/db/models/users"
	"github.com/meedeley/go-launch-starter-code/internal/conf"
	"github.com/meedeley/go-launch-starter-code/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	db *pgxpool.Pool
}

func NewUserUseCase(db *pgxpool.Pool) UserUseCase {
	return UserUseCase{db: db}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (u *UserUseCase) Register(ctx context.Context, userReq entity.UserRegisterRequest) (*entity.UserRegisterResponse, error) {
	db := u.db
	q := users.New(db)

	hashedPassword, _ := hashPassword(userReq.Password)

	row, err := q.InsertUser(ctx, users.InsertUserParams{
		Name:     userReq.Name,
		Email:    userReq.Email,
		Password: hashedPassword,
	})

	var updatedAt *time.Time
	if row.UpdatedAt.Valid {
		updatedAt = &row.UpdatedAt.Time
	}

	userRes := entity.UserRegisterResponse{
		Id:        row.ID,
		Name:      row.Name,
		Email:     row.Email,
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: updatedAt,
	}

	return &userRes, err
}

func (u UserUseCase) Login(ctx context.Context, userReq entity.UserLoginRequest) (*entity.UserLoginResponse, error) {
	db := u.db
	defer db.Close()
	q := users.New(db)

	email := userReq.Email
	pass := userReq.Password

	result, err := q.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	checked := CheckPasswordHash(pass, result.Password)
	if !checked {
		return nil, err
	}

	token := jwt.New(jwt.SigningMethodHS256)

	var claims jwt.MapClaims = token.Claims.(jwt.MapClaims)
	claims["user_id"] = result.ID
	claims["email"] = result.Email
	claims["name"] = result.Name
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte(conf.JwtSecret()))
	if err != nil {
		return nil, err
	}

	userRes := entity.UserLoginResponse{
		Id:        result.ID,
		Name:      result.Name,
		Email:     result.Email,
		Token:     tokenString,
		CreatedAt: result.CreatedAt.Time,
		UpdatedAt: result.UpdatedAt.Time,
	}

	return &userRes, nil
}
