package models

import "gorm.io/gorm"

// Vår notifications table
type Notification struct {
	gorm.Model
	UserId uint // Koppling till användarens id
	Message string
	IsRead bool

}

