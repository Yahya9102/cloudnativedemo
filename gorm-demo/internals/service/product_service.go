package service // logiklagret

import (
	"cloudnativedemo/gorm-demo/internals/models"
	"cloudnativedemo/gorm-demo/internals/repository"
)

// ProductService innehåller affärslogiken
type ProductService struct {
	repo *repository.ProductRepository
}

// Skapa ny service med repository som beroende
func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

// Skapa produkt
func (s *ProductService) CreateProduct(name string, price int) models.Product {
	product := models.Product{Name: name, Price: price}
	return s.repo.Create(product)
}

// Hämta alla produkter
func (s *ProductService) ListProducts() []models.Product {
	return s.repo.GetAll()
}

// Uppdatera produkt
func (s *ProductService) UpdateProduct(id uint, name string, price int) (models.Product, bool) {
	return s.repo.Update(id, name, price)
}

// Radera produkt
func (s *ProductService) DeleteProduct(id uint) bool {
	return s.repo.Delete(id)
}
