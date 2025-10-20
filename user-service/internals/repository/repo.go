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
	dsn := "root:12345@tcp(127.0.0.1:3306)/users?charset=utf8mb4&parseTime=True&loc=Local"

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




/*




func (r *UserRepository) GetAll() []models.User{
	return r.users
}




func (r *UserRepository) Delete(id int) bool{
	
	for i, u := range r.users {
		if u.ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return  true
		}
	}
	return false
}



func(r *UserRepository) Update(id int, updated models.User) (models.User, bool) {


	// Loopar genom alla våra användare i vår slice
	for i, u := range r.users {
		if u.ID == id { // Om ID matchar
			r.users[i].Name = updated.Name // Updatera namn
			r.users[i].Age = updated.Age // Updatera age
			

			// Returnera uppdaterade användare + true
			return r.users[i], true
		}
	}

	// or else inge användare hittades med det ID:t
	return models.User{}, false
}


*/