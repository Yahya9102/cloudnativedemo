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


