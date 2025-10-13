package service

import (
	"cloudnativedemo/internals/models"
	"cloudnativedemo/internals/repository"
)

//UserService ansvarar för logik kring användare
type UserService struct {
	repo *repository.UserRepository 
}


// Skapar en userService
func NewUserService(repo *repository.UserRepository) *UserService{
	return &UserService{repo: repo}
}


func (s *UserService) CreateUser(name string, age int) models.User {
	user := models.User{Name: name, Age: age}
	return s.repo.Add(user)
}

func (s *UserService) ListUsers() []models.User{
	return s.repo.GetAll()
}

func (s *UserService) DeleteUser(id int) bool {
	return s.repo.Delete(id)
}


func (s *UserService) UpdateUsers(id int, name string, age int) (models.User, bool){
	// skapa ny användare med uppdaterad meterial
	updated := models.User{Name: name, Age: age}

	// kalla på repo funktion och mata in objektet
	return s.repo.Update(id, updated)
}
