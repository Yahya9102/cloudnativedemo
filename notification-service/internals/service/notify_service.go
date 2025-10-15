package service

import (
	"cloudnativedemo/notification-service/internals/client"
	"cloudnativedemo/notification-service/internals/models"
	"fmt"
)



type NotifyService struct {
	userClient *client.UserClient
	logs [] string // Vi lagrar en pekare till vår userClient
}


// Den skapar en ny instans av NotifyService
func NewNotifyService(userClient * client.UserClient) *NotifyService{
	return &NotifyService{
		userClient: userClient,
		logs: []string{},
	}
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
	s.logs = append(s.logs, logs...)
	return logs

}




// Funktion för att filtrera användare baserat på deras ålder 18+

func (s *NotifyService) FetchAdultUsers() ([]models.User, error){
	users, err := s.userClient.FetchAllUsers();
	
	if err != nil {
		return nil, err
	}

	var adults []models.User

	for _, user := range users {
		if user.Age >= 18{
			adults = append(adults, user)
		}
	}
	return adults, nil
}


// Return all logs
func (s *NotifyService) GetLogs() []string {
	return s.logs
}


// Clear logs
func (s *NotifyService) ClearLogs(){
	s.logs = []string{}
}


// send notification to a specific user via ID
func (s *NotifyService) SendNotificationToUser(id int, message string) (string, error) {

	// Hämta alla användare

	users, err := s.userClient.FetchAllUsers()

	if err != nil {
		return "", fmt.Errorf("couldn´t fetch user %v", err.Error())
	}

	// Loopa genom användarna

	for _, user := range users {
		if user.ID == id {
			log := fmt.Sprintf("Notis skickad till %s (ID %d): %s", user.Name, user.ID, message)
			s.logs = append(s.logs, log)
			return log, nil
		
		}
	}

	return "", fmt.Errorf("inge användare med id %d hittades", id)

}