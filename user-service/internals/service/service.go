package service

import (
	"cloudnativedemo/user-service/internals/models"
	"cloudnativedemo/user-service/internals/repository"
)

//UserService ansvarar för logik kring användare
type UserService struct {
	repo *repository.UserRepository 
}


// Skapar en userService
func NewUserService(repo *repository.UserRepository) *UserService{
	return &UserService{repo: repo}
}


func (s *UserService) CreateUser(name string, age int) (*models.User, error) {
	user := &models.User{Name: name, Age: age}
	err:= s.repo.CreateUser(user)
	return user, err

}


func (s *UserService) ListUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}




func (s *UserService) UpdateUser(id uint, name string, age int) error{
	return s.repo.UpdateUser(id, name, age)	
}


func (s *UserService) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}

