package repo_sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	entity2 "mirgalievaal-project/backend/internal/entity"
)

func NewSQLIte(dbUri string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbUri), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&entity2.User{}, &entity2.Product{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
