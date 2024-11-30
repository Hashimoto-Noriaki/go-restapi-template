package repositories

import (
	"errors"
	"go-restapi-boiltertemplate/app/models"

	"gorm.io/gorm"
)

type PAuthRepository interface {
	CreateUser(user models.User) error
	FindUser(email string) (*models.User, error)
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

func (r *AuthRepository) FindUser(email string) (*models.User, error){
	var user models.User
	result := r.db.First(&user, "email = ?", email)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, errors.New("ユーザーが見つかりません")
		}
		return nil, result.Error
	}
	return &user, nil
} 
