package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"not null;type:varchar(255)"`
	Email     string `json:"email" gorm:"not null;type:varchar(255)"`
	Password  string `json:"password" gorm:"not null;type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
