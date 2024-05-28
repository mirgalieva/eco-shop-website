package entity

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAd   gorm.DeletedAt `gorm:"index"`
	Description string
	Name        string
	Price       string
	SellerID    uint
	Tag         string
}
type ProductRepository interface {
	Create(seller *Product) error
	GetAll() (*[]Product, error)
	Get(id uint) (*Product, error)
	Update(user *Product) error
	Delete(id uint) error

	GetUserByID(userID uint) (*User, error)
	//GetUserByID(productID uint) *entity.User
}

type ProductService interface {
	Create(product *Product) (*Product, error)
	Get(id uint) (*Product, error)
	GetAll() (*[]Product, error)
	Update(product *Product) error
	Delete(product *Product) error
}
