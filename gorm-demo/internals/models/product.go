package models // package för databasmodeller

import "gorm.io/gorm"

// Product representerar tabellen "products" i MySQL-databasen
type Product struct {
	gorm.Model        // lägger till ID, CreatedAt, UpdatedAt, DeletedAt automatiskt
	Name  string `json:"name"`  // produktens namn
	Price int    `json:"price"` // produktens pris
}
