package service

import (
	"cloudnativedemo/notification-service/internals/client"
	"cloudnativedemo/notification-service/internals/models"
	"fmt"
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


// Skapa en ny funktion som skickar notis till användare
func (s *NotifyService) SendNotification(message string)[] string {


	var logs []string

	users, err := s.userClient.FetchAllUsers()

	if err != nil {
		logs = append(logs, fmt.Sprintf("kunde inte nå user-service: %v", err.Error()))
		return logs
	}

	// Gå genom varje användare (aka loopa genom och "skicka" en notis)
	for _, user := range users {
		log := fmt.Sprintf("Skickade notis till %s: %s", user.Name, message)
		logs = append(logs, log)
		fmt.Println(log)
	}
	return logs

}