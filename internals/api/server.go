package api

import (
	"cloudnativedemo/internals/service"
	"encoding/json"
	"net/http"
)

// Starta http server som kan hantera requests
func StartServer(userService *service.UserService){


	// registrerar en route för vår endpoint /users
	http.HandleFunc("/users", func(w http.ResponseWriter, r * http.Request) {

		// Kontroller om metod = get
		if r.Method == http.MethodGet {

			// Hämta alla våra användare
			users := userService.ListUsers()

			// S ätta svaret content type till JSON
			w.Header().Set("Content-Type", "application/json")

			// Skicka tillbaka alla användare som JSON
			json.NewEncoder(w).Encode(users)


		} else {
			
			// Returnera error om metod inte är get
			http.Error(w, "Metoden är inte tillåten ", http.StatusMethodNotAllowed)
		}

	})

	// Starta servern på port 8080
	http.ListenAndServe(":8080", nil)

}