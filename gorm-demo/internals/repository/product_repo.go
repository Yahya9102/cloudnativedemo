package repository // ansvarar för databasåtkomst

import (
	"cloudnativedemo/gorm-demo/internals/models" // vår modell
	"fmt"

	"gorm.io/driver/mysql" // MySQL-driver
	"gorm.io/gorm"
)

// ProductRepository håller i databaslogik
type ProductRepository struct {
	db *gorm.DB
}

// NewProductRepository ansluter till MySQL och migrerar modellen
func NewProductRepository() *ProductRepository {
	// Format för MySQL-anslutning
	// dsn = Data Source Name
	dsn := "root:12345@tcp(127.0.0.1:3306)/gorm_demo?charset=utf8mb4&parseTime=True&loc=Local"

	// Anslut till databasen
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Kunde inte ansluta till MySQL: %v", err))
	}

	// Skapa tabeller om de inte finns
	db.AutoMigrate(&models.Product{})

	// Returnera en instans av repository
	return &ProductRepository{db: db}
}

// Create lägger till ny produkt
func (r *ProductRepository) Create(product models.Product) models.Product {
	r.db.Create(&product)
	return product
}

// GetAll hämtar alla produkter
func (r *ProductRepository) GetAll() []models.Product {
	var products []models.Product
	r.db.Find(&products)
	return products
}

// Update uppdaterar produkt baserat på ID
func (r *ProductRepository) Update(id uint, name string, price int) (models.Product, bool) {
	var product models.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return models.Product{}, false
	}
	product.Name = name
	product.Price = price
	r.db.Save(&product)
	return product, true
}

// Delete tar bort produkt baserat på ID
func (r *ProductRepository) Delete(id uint) bool {
	var product models.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return false
	}
	r.db.Delete(&product)
	return true
}
