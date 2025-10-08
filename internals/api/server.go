package api

import (
	"cloudnativedemo/internals/service"
	"encoding/json"
	"net/http"
)

// Starta http server som kan hantera requests
func StartServer(userService *service.UserService){


	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodGet:
			users := userService.ListUsers()

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(users)


		case http.MethodPost:
			
			// Skapa json bodyn som ska skickas in
			var input struct{	
				Name string `json:"name"`
				Age int `json:"age"` 
			}	

				// Avkodar och error check
			if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
				http.Error(w, "Icke JSON", http.StatusBadRequest)
				return
			}

			// Skapa ny user
			user := userService.CreateUser(input.Name, input.Age)

			// Returnera http status kod
			w.WriteHeader(http.StatusCreated)

			// returnera skapade användare
			json.NewEncoder(w).Encode(user)
		
		} 

	})

	http.ListenAndServe(":8080", nil)

}