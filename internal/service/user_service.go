package service

import (
	"context"
	"time"

	"go-user-api/db/sqlc"
	"go-user-api/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) Repo() *repository.UserRepository {
	return s.repo
}

// Parse DOB string and create user
func (s *UserService) CreateUser(ctx context.Context, name, dobStr string) (sqlc.User, error) {
	dob, err := time.Parse("2006-01-02", dobStr)
	if err != nil {
		return sqlc.User{}, err
	}
	return s.repo.CreateUser(ctx, name, dob)
}

// Age calculation
func CalculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()

	if now.YearDay() < dob.YearDay() {
		age--
	}
	return age
}
