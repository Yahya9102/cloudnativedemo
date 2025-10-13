package api

import (
	"cloudnativedemo/user-service/internals/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// funktion som start Gin server
func StartServer(userService *service.UserService) {

	
	// Skapar en default server med logger och recovery	
	r := gin.Default()

	// Vår Get request där vi hämtar alla användare
	r.GET("/users", func(c *gin.Context) {
		users := userService.ListUsers()
		c.JSON(http.StatusOK, users)
	})



	// Vår post anrop
	r.POST("/users", func(c *gin.Context) {
		// Skapa en kompia av struct för att läsa in JSON från request
		var input struct{	
			Name string `json:"name"`
			Age int `json:"age"` 
		}

		// validerat formatet
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		} 

		// skapa user via service lager
		user := userService.CreateUser(input.Name,input.Age)

		// returnerar vi en 201, med user details
		c.JSON(http.StatusOK, user)
	})


	r.DELETE("/users/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		var id int
		fmt. Sscanf(idStr, "%d", &id) // Konvertera från string to int

		// User deleted if found
		if userService.DeleteUser(id) {
			c.Status(http.StatusNoContent)
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		}

	})


	r.PUT("/users/:id", func(c *gin.Context) {

		idStr := c.Param("id")
		var id int
		fmt. Sscanf(idStr, "%d", &id) // Konvertera från string to int


		// Skapa en kompia av struct för att läsa in JSON från request
		var input struct{	
			Name string `json:"name"`
			Age int `json:"age"` 
		}

		// validerat formatet
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		} 

		user, ok := userService.UpdateUsers(id, input.Name, input.Age)

		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		}

		c.JSON(http.StatusOK, user)

	})
	
	r.Run(":8080")

}






/*
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

		case http.MethodDelete:
			// Hämta id från url users?id=1
			idStr := r.URL.Query().Get("id")
			var id int
			fmt. Sscanf(idStr, "%d", &id) // Konvertera från string to int

			fmt.Println(id)
			if userService.DeleteUser(id) {
				w.WriteHeader(http.StatusNoContent)
			} else {
				http.Error(w, "User not found", http.StatusNotFound)
			}



		case http.MethodPut:
			// Hämta id från url users?id=1
			idStr := r.URL.Query().Get("id")
			var id int
			fmt. Sscanf(idStr, "%d", &id) // Konvertera från string to int
		
			var input struct{	
				Name string `json:"name"`
				Age int `json:"age"` 
			}

			// Avkodar och error check
			if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
				http.Error(w, "Icke JSON", http.StatusBadRequest)
				return
			}
			
			// Updatera vår användare
			user, ok := userService.UpdateUsers(id,input.Name, input.Age)

			// Vid inte okej request
			if !ok {
				http.Error(w, "Användare hittades ej", http.StatusNotFound)
				return
			}
			// returnera uppdaterade användaren
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(user)
		default: 
			http.Error(w, "Something went wrong!", http.StatusBadRequest)
		
			} 
	})
	// Starta server på port 8080
	http.ListenAndServe(":8080", nil)

}


*/