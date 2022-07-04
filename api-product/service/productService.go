package service

import (
	"test-azuralabs/api-product/models"
	"test-azuralabs/api-product/repository"
)

type service struct {
	repository repository.Repository
}

type Service interface {
	GetAllProduct() ([]models.Product, error)
	GetProductByID(id int) (models.Product, error)
	CreateProduct(productReq models.Product) (models.Product, error)
}

func NewService(repositoey repository.Repository) *service {
	return &service{repository: repositoey}
}

func (s *service) GetAllProduct() ([]models.Product, error) {
	products, err := s.repository.GetAllProduct()
	return products, err
}

func (s *service) GetProductByID(id int) (models.Product, error) {
	product, err := s.repository.GetProductByID(id)
	return product, err
}

func (s *service) CreateProduct(productReq models.Product) (models.Product, error) {
	product, err := s.repository.CreateProduct(productReq)
	return product, err
}
