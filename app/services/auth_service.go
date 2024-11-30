package services

import (
	"go-restapi-boiltertemplate/app/models"
	"go-restapi-boiltertemplate/app/repositories"

	"golang.org/x/crypto/bcrypt"
)

type PAuthService interface {
	Signup(email string, password string) error 
}

type AuthService struct {
	repository repositories.PAuthRepository
}

func NewAuthService(repository repositories.PAuthRepository) PAuthService {
	return &AuthService{repository: repository}
}

func (s *AuthService) Signup(email string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		Email: email,
		Password: string(hashedPassword),
	}
	return s.repository.CreateUser(user)
}