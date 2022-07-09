package repository

import (
	"github.com/marcos-nsantos/e-commerce/auth-service/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Create(user *model.User) error {
	return ur.db.Create(user).Error
}
