package repository

import (
	"test-azuralabs/api-product/models"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

type Repository interface {
	GetAllProduct() ([]models.Product, error)
	GetProductByID(id int) (models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
}

func NewRepository(DB *gorm.DB) *repository {
	return &repository{DB: DB}
}

func (r *repository) GetAllProduct() ([]models.Product, error) {
	var products []models.Product
	err := r.DB.Find(&products).Error
	return products, err
}

func (r *repository) GetProductByID(id int) (models.Product, error) {
	var product models.Product
	err := r.DB.Find(&product, id).Error
	return product, err
}

func (r *repository) CreateProduct(product models.Product) (models.Product, error) {
	err := r.DB.Create(&product).Error
	return product, err
}
