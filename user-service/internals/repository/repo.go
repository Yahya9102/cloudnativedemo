package repository

import (
	"cloudnativedemo/user-service/internals/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// UserRepo  ansvarar för DB operationer
type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() *UserRepository {
	

/*

	// DSN (Data Source Name) för PostgreSQL
	dsn := "host=127.0.0.1 user=admin password=admin123 dbname=accounts port=3307 sslmode=disable TimeZone=Europe/Stockholm"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Kunde inte ansluta till databasen: %v", err)
	}

*/



	// Definera DSN (Data Source Name)
	dsn := "root:12345@tcp(127.0.0.1:3306)/accounts?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Kunde inte ansluta till databasen: %v", err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Notification{})

	if err != nil {
		log.Fatalf("Kunde inte migrera modellen: %v", err)
	}

	log.Println("Ansluten till databasen :)")


	return &UserRepository{DB: db}

}


func(r * UserRepository) GetNotificationsByUser(userID uint) ([]models.Notification, error) {
	var notifications []models.Notification

	err := r.DB.Where("user_id = ?", userID).Find(&notifications).Error
	return notifications, err
}


func (r *UserRepository) CreateNotification(notification *models.Notification) error {
	return r.DB.Create(notification).Error
}


func(r *UserRepository) UpdateNotification(id uint, isRead bool) error {
	return r.DB.Model(&models.Notification{}).Where("id = ?", id).Update("is_read",isRead).Error
}


func(r *UserRepository) DeleteNotification(id uint) error {
	return r.DB.Unscoped().Delete(&models.Notification{}, id).Error
}




func (r *UserRepository) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}


func (r *UserRepository) GetAllUsers() ([]models.User, error){
	var users [] models.User
	err := r.DB.Find(&users).Error
	return users, err
	
}


func(r *UserRepository) UpdateUser(id uint, name string, age int) error {
	return r.DB.Model(&models.User{}).Where("id = ?", id).Updates(models.User{Name: name, Age: age}).Error
}


func (r *UserRepository) DeleteUser(id uint) error{
	return r.DB.Unscoped().Delete(&models.User{}, id).Error
}

