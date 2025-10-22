package repository

import (
	"cloudnativedemo/notification-service/internals/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}


func NewRepository() *Repository {

	dns := "root:12345@tcp(127.0.0.1:3306)/notifications?charset=utf8mb4&parseTime=True&loc=Local"


	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatalf("Kunde inte ansluta till database: %v", err)
	}

	if err := db.AutoMigrate(&models.Notification{}); err != nil {
		log.Fatalf("Kunde inte migrera notifications modellen: %v", err)	
	}

	log.Println("Ansluten till notifications databasen")

	return &Repository{DB: db}

}


func(r *Repository) CreateNotification(notification *models.Notification) error {
	return r.DB.Create(notification).Error
} 

func(r *Repository) GetAllNotifications() ([]models.Notification, error) {
	var notifications []models.Notification

	err := r.DB.Find(&notifications).Error

	return notifications, err
}