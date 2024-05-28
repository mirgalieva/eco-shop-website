package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAd gorm.DeletedAt `gorm:"index"`
	UserRegister
}

type UserLogin struct {
	Email    string `gorm:"unique"`
	Password string `json:"Password,omitempty"`
}

type UserRegister struct {
	UserLogin

	FirstName string
	LastName  string
}
type UserRepository interface {
	Create(seller *User) error
	GetAll() (*[]User, error)
	Get(id uint) (*User, error)
	Update(user *User) error
	Delete(id uint) error
	GetByEmail(email string) (*User, error)
}

type UserService interface {
	Get(id uint) (*User, error)
	GetAll() (*[]User, error)
	Update(user *User) error
	Delete(user *User) error

	Register(userRegister *UserRegister) error
	Login(userLogin *UserLogin) error
}

func (u *User) OmitPassword() {
	u.Password = ""
}
