package service

import (
	"cloudnativedemo/gorm-demo/internals/models"
	"cloudnativedemo/gorm-demo/internals/repository"
)

// ProductService hanterar affärslogiken för produkter
type ProductService struct {
	repo *repository.ProductRepository
}

// NewProductService skapar ny service och kopplar till repiotory
func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}


// CreateProduct skapar ny produckt via repository
func (s *ProductService) CreateProduct(name string, price int) models.Product {
	product := models.Product{Name: name, Price: price}
	return s.repo.Add(product)
}


// GetAllProducts hämtar alla produkter via repository
func(s *ProductService) ListProducts() []models.Product{
	return s.repo.GetAll()
}