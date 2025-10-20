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
	
	// Definera DSN (Data Source Name)
	dsn := "root:12345@tcp(127.0.0.1:3306)/accounts?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Kunde inte ansluta till databasen: %v", err)
	}

	err = db.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatalf("Kunde inte migrera modellen: %v", err)
	}

	log.Println("Ansluten till databasen :)")


	return &UserRepository{DB: db}

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

