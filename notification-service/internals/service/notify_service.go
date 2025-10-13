package service

import (
	"cloudnativedemo/notification-service/internals/client"
	"cloudnativedemo/notification-service/internals/models"
)



type NotifyService struct {
	userClient *client.UserClient // Vi lagrar en pekare till vår userClient
}


// Den skapar en ny instans av NotifyService
func NewNotifyService(userClient * client.UserClient) *NotifyService{
	return &NotifyService{userClient: userClient}
}


// Vi hämtar alla users från user-service
func (s *NotifyService) FetchUsers() ([]models.User, error) {
	users, err := s.userClient.FetchAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}
