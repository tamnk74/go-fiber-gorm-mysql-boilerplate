package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User type
type User struct {
	ID        uint           `gorm:"primary_key" json:"id"`
	Name      string         `gorm:"type:varchar(255);unique,not null" json:"name"`
	Email     string         `gorm:"type:varchar(100);unique;not null;unique_index" json:"email"`
	Password  string         `json:"-"`
	Status    uint           `json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	u.Password = string(hashedPassword)
	return
}
