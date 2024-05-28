package repository

import (
	"gorm.io/gorm"
	entity2 "mirgalievaal-project/backend/internal/entity"
	repo_sqlite "mirgalievaal-project/backend/internal/repository/sqlite"
)

type Repository struct {
	entity2.UserRepository
	entity2.ProductRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository:    repo_sqlite.NewUserSQLite(db),
		ProductRepository: repo_sqlite.NewProductSQLite(db),
	}
}
