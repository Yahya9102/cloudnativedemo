package repository

import (
	"cloudnativedemo/internals/models"
)

type UserRepository struct {
	users  []models.User
	nextID int
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users:  []models.User{},
		nextID: 1,
	}
}


func (r *UserRepository) Add(user models.User) models.User {
	user.ID = r.nextID
	r.nextID++
	r.users = append(r.users, user)
	return user
}


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


