package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey;->"`
	Name      string         `json:"name" gorm:"not null;type:varchar(255)" validate:"required,notblank,min=3,max=255"`
	Email     string         `json:"email" gorm:"not null;type:varchar(255)" validate:"required,notblank,email"`
	Password  string         `json:"password,omitempty" gorm:"not null;type:varchar(64)" validate:"required,notblank,min=8,max=64"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
