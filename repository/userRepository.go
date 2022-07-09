package repository

import (
	"github.com/marcos-nsantos/e-commerce/auth-service/model"
	"gorm.io/gorm"
)

type userAPI struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Create(user *model.User) error {
	return ur.db.Create(user).Error
}

func (ur *UserRepository) FindByID(id uint) (*userAPI, error) {
	var user userAPI
	err := ur.db.Model(&model.User{}).First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) Update(user *model.User) error {
	return ur.db.Model(&user).Select("name, email").Updates(user).Error
}

func (ur *UserRepository) UpdatePassword(user *model.User) error {
	return ur.db.Model(&user).Update("password", user.Password).Error
}
