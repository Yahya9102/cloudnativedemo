package client

import (
	"cloudnativedemo/notification-service/internals/models"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserClient struct {
	BaseURL string
}

// NewUserClient skapar en ny instans av UserClient
func NewUserClient(baseURl string) *UserClient {
	return &UserClient{
		BaseURL: baseURl,
	}
}

func (c *UserClient) FetchAllUsers() ([]models.User, error) {
	// Här skulle du implementera logiken för att anropa user-service

	// Vår URL tex localhost:8808/	
	url := c.BaseURL + "/users"


	// Gör ett Get anrop mot vår user-service
	resp , err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("kunde inte nå user-service: %v", err)
	}
	
	// Stäng anslutningen när vi är klara!
	defer resp.Body.Close()


	// Kontrollera att vi fick en 200 OK
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("felaktigt svar från user-service: %v", resp.Status)
	}


	 // Tolka JSON svaret till en SLICE AV USERS

	 var users []models.User
	 if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, fmt.Errorf("kunde inte tolka JSON ")
	 }

	 return users, nil


}
