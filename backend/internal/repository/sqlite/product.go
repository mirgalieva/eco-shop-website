package repo_sqlite

import (
	"errors"
	"gorm.io/gorm"
	entity2 "mirgalievaal-project/backend/internal/entity"
)

type ProductSQLite struct {
	db *gorm.DB
}

func (r *ProductSQLite) GetUserByID(userID uint) (*entity2.User, error) {
	user := entity2.User{}
	result := r.db.Where("id = ?", userID).Last(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, entity2.ErrUserNotFound
	} else if result.Error != nil {
		return nil, result.Error
	} else {
		return &user, nil
	}
}

func NewProductSQLite(db *gorm.DB) *ProductSQLite {
	return &ProductSQLite{db: db}
}
func (r *ProductSQLite) GetAll() (*[]entity2.Product, error) {
	var products []entity2.Product

	if result := r.db.Find(&products); result.Error != nil {
		return nil, result.Error
	}
	return &products, nil
}

func (r *ProductSQLite) Get(id uint) (*entity2.Product, error) {
	var product entity2.Product
	if result := r.db.Where("id = &", id).First(&product); result.Error == nil {
		return &product, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &product, entity2.ErrProductNotFound
	} else {
		return &product, result.Error
	}
}

func (r *ProductSQLite) Create(product *entity2.Product) error {
	if result := r.db.Create(product); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (r *ProductSQLite) Update(product *entity2.Product) error {
	result := r.db.Model(product).Updates(product)
	if result != nil {
		return result.Error
	} else {
		return nil
	}
}

func (r *ProductSQLite) Delete(id uint) error {
	result := r.db.Delete(&entity2.Product{}, id)
	if result != nil {
		return result.Error
	} else {
		return nil
	}
}
