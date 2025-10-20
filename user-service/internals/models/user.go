package models

import "gorm.io/gorm"

type User struct {
	gorm.Model //Auto inkluderar ID, CreatedAt, UpdatedAt, DeletedAt
	Name string
	Age  int
}