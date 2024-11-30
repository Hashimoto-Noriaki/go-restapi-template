package repositories

import (
	"go-restapi-boiltertemplate/app/models"

	"gorm.io/gorm"
)

type PAuthRepository interface {
	CreateUser(user models.User) error
}

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB)  PAuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository)  CreateUser(user models.User) error {
	result :=  r.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}