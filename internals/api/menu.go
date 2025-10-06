package api

import (
	"fmt"
)


func add(a, b float64) float64 {
	return  a + b
}



func RunMenu (){

	for {
		fmt.Println("1: add")
		fmt.Println("4: Avsluta")


		var choice int;
		fmt.Println("Välj: ")

		fmt.Scanln(&choice)

		switch choice {
		case 1: 
			 var numberOne, numberTwo float64;
		
			fmt.Println("Skriv in första talet")
			fmt.Scanln(&numberOne)
			
			fmt.Println("Skriv in andra talet..")
			fmt.Scanln(&numberTwo)

			
			result := add(numberOne, numberTwo)
			fmt.Printf("Resultat: %.2f \n", result)

		case 2:
			fmt.Println("Avslutar programmet")	
			return
	
		default: 
		fmt.Println("Ogiltig val försök igen please...")
		}
		
	}


	
}