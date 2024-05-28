package service

import (
	entity2 "mirgalievaal-project/backend/internal/entity"
)

type ProductService struct {
	productRepo entity2.ProductRepository
	userService entity2.UserService
}

func (s *ProductService) GetAll() (*[]entity2.Product, error) {
	productDB, err := s.productRepo.GetAll()
	if err != nil {
		return nil, err
	} else {
		return productDB, nil
	}
}

func NewProductService(productRepo entity2.ProductRepository, userService UserService) *ProductService {
	return &ProductService{
		productRepo: productRepo,
		userService: &userService,
	}
}
func (s *ProductService) Create(product *entity2.Product) (*entity2.Product, error) {
	/*_, err := s.productRepo.GetUserByID(product.SellerID)
	if err != nil {
		return nil, err
	} else {*/
	err := s.productRepo.Create(product)
	if err != nil {
		return nil, err
	} else {
		return product, nil
	}
} /*
}*/
func (s *ProductService) Get(id uint) (*entity2.Product, error) {
	productDB, err := s.productRepo.Get(id)
	if err != nil {
		return nil, err
	} else {
		return productDB, nil
	}
}
func (s *ProductService) Update(product *entity2.Product) error {
	_, err := s.productRepo.Get(product.ID)
	if err != nil {
		return err
	}
	err = s.productRepo.Update(product)
	return err
}

func (s *ProductService) Delete(product *entity2.Product) error {
	err := s.productRepo.Delete(product.ID)
	return err
}
