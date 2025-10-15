package models

import "gorm.io/gorm"

// Product representerar en produkt i vår databas
type Product struct {
	gorm.Model // Lägger automatiskt till ID, createdAt, updatedAt, deletedAt
	Name string
	Price int
}