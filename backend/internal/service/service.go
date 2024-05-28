package service

import (
	entity2 "mirgalievaal-project/backend/internal/entity"
	"mirgalievaal-project/backend/internal/repository"
)

type Service struct {
	User    entity2.UserService
	Product entity2.ProductService
}

func NewService(repo *repository.Repository) *Service {
	userService := NewUserService(repo.UserRepository)
	return &Service{
		User:    userService,
		Product: NewProductService(repo.ProductRepository, *userService),
	}
}
