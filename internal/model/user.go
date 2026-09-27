package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"type:varchar(100);not null"`
	Email     string         `json:"email" gorm:"type:varchar(100);uniqueIndex;not null"`
	Password  string         `json:"-" gorm:"not null"` // json:"-" = hidden dari response
	Role      string         `json:"role" gorm:"type:varchar(20);default:buyer"`
	Phone     string         `json:"phone" gorm:"type:varchar(20)"`
	Address   string         `json:"address" gorm:"type:text"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (u *User) HashPassword(plain string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashed)
	return nil
}

func (u *User) CheckPassword(plain string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain))
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	Phone string `json:"phone"`
}

func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
		Role:  u.Role,
		Phone: u.Phone,
	}
}
