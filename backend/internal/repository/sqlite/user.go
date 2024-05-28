package repo_sqlite

import (
	"errors"
	"gorm.io/gorm"
	entity2 "mirgalievaal-project/backend/internal/entity"
)

type UserSQLite struct {
	db *gorm.DB
}

func NewUserSQLite(db *gorm.DB) *UserSQLite {
	return &UserSQLite{db: db}
}
func (r *UserSQLite) GetAll() (*[]entity2.User, error) {
	var users []entity2.User

	if result := r.db.Find(&users); result.Error != nil {
		return nil, result.Error
	}
	return &users, nil
}

func (r *UserSQLite) Get(id uint) (*entity2.User, error) {
	var user entity2.User
	if result := r.db.Where("id = ?", id).First(&user); result.Error == nil {
		return &user, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &user, entity2.ErrUserNotFound
	} else {
		return &user, result.Error
	}
}

func (r *UserSQLite) GetByEmail(email string) (*entity2.User, error) {
	var user entity2.User
	if result := r.db.Where("email = ?", email).First(&user); result.Error == nil {
		return &user, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &user, entity2.ErrUserNotFound
	} else {
		return &user, result.Error
	}
}

func (r *UserSQLite) Create(user *entity2.User) error {
	if result := r.db.Create(user); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (r *UserSQLite) Update(user *entity2.User) error {
	result := r.db.Model(user).Updates(user)
	if result != nil {
		return result.Error
	} else {
		return nil
	}
}

func (r *UserSQLite) Delete(id uint) error {
	result := r.db.Delete(&entity2.User{}, id)
	if result != nil {
		return result.Error
	} else {
		return nil
	}
}
