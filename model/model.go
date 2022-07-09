package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey;->"`
	Name      string `json:"name" gorm:"not null;type:varchar(255)"`
	Email     string `json:"email" gorm:"not null;type:varchar(255)"`
	Password  string `json:"password" gorm:"not null;type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
