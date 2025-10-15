package repository

import (
	"cloudnativedemo/gorm-demo/internals/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB // Pekare till aktiva databasanslutningen
}

func NewProductRepository() *ProductRepository {

	dsn := "root:12345@tcp(127.0.0.1:3306)/gorm_demo?charset=utf8mb4&parseTime=True&loc=Local" // Databas url

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // Öppnar anslutning till mysql
	if err != nil {
		log.Fatalf("Kunde inte ansluta till databasen %v", err)
	}

	db.AutoMigrate(&models.Product{})  // Skapar tabell i databasen om den inte finns
	return &ProductRepository{db: db} //Returnera en pekare till repository structen
}

// Add sparar en ny produkt i databasen
func (r *ProductRepository) Add(product models.Product) models.Product{
	r.db.Create(&product) // Skapar en ny produkt i databasen
	return product
}


// GetAll hämtar alla produkter från databasen
func (r *ProductRepository) GetAll() []models.Product{
	var products []models.Product // Slice för att lagra alla produkter
	r.db.Find(&products)  // Select * from products
	return products 
}

