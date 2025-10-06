package api

import (
	"cloudnativedemo/internals/service"
	"fmt"
)


func RunMenu (userService *service.UserService){

	for {
		fmt.Println("1: Skapa användare")
		fmt.Println("2: Lista alla användare")
		fmt.Println("3: Avsluta")


		var choice int;
		fmt.Println("Välj: ")

		fmt.Scanln(&choice)

		switch choice {
		case 1: 
			var name string
			var age int
		
			fmt.Println("Skriv in ditt namn..")
			fmt.Scanln(&name)
			
			fmt.Println("Skriv in ålder..")
			fmt.Scanln(&age)

			userService.CreateUser(name,age)
			

		case 2:
			users := userService.ListUsers()
			if len(users) == 0 {
				fmt.Println("Finns inga användare")
			} else {
				fmt.Println("Här är alla våra användare")
				for _, u := range users {
					fmt.Printf("ID: %d, name: %s, age: %d \n", u.ID, u.Name, u.Age)
				}
			}
			 
		case 3:
			fmt.Println("Avslutar programmet")	
			return

		default: 
		fmt.Println("Ogiltig val försök igen please...")
		}
		
	}


	
}