package main

import (
	"fmt"
)

// Deklarera med var utanför funktioner, alternativ med rätt typ
var name = "Yahya"
var age = 34

// Med typer
var username string = "yahya"
var userAge int8 = 34;

func main(){

	// I funktioner

	firstname := "Yahya"
	randomNumber := 3

	// Med formatsträng
	fmt.Printf("Hello! %s, your number is %d\n", firstname, randomNumber)

	// Eller enklare med Println
	fmt.Println("Hello!", firstname, "your number is", randomNumber)


	fmt.Println(username, userAge)
	fmt.Println(name, age)
}