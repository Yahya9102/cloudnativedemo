package main

import (
	"cloudnativedemo/user-service/internals/api"
	"cloudnativedemo/user-service/internals/repository"
	"cloudnativedemo/user-service/internals/service"
)



func main(){
	 repo := repository.NewUserRepository()
	 userService := service.NewUserService(repo)
	// api.RunMenu(userService)
	api.StartServer(userService)

}

/*
	Day 3
	Struckts 

	
type User struct {
	Name  string
	Email string
}


func (u *User) Greet() {
	fmt.Println("Hej:", u.Name)
}



func main() {
	user := User{Name: "Yahya", Email: "yahya.hussein@systementor.se"}

	user.Greet()

	

	user.Name = "Anton";

	user.Greet()


}



*/



/*Day 2 - Loops, Arrays och Slices, funktioner och mappar




// SOLID principles

	// 1. Single Responsibility Principle

	// 2. Open/Closed Principle

	// 3. Liskov Substitution Principle

	// 4. Interface Segregation Principle

	// Djur kan fly(), simma(), springa()

	// 5. Dependency Inversion Principle

	// Swedbank

	// Microservices

	// Auth {CONTAINERS 1} {CONTAINERS 2} {CONTAINERS 3}

	// Valutakonvertering

	// Aktieköp {CONTAINERS 1} {CONTAINERS 2} {CONTAINERS 3}

	// banköverföring




	func returnFunc(x string) func(){
		return func ()  {
			fmt.Println(x)
		}
	}


	returnFunc("Hello")()
	functionValue := returnFunc("Goodbye")
	functionValue()



	 names := []string{}

    var name1, name2 string
    fmt.Print("Skriv första namnet: ")
    fmt.Scanln(&name1)
    fmt.Print("Skriv andra namnet: ")
    fmt.Scanln(&name2)

    names = append(names, name1, name2)

    printNames(names)



	func printNames(slice []string) {
		for _, name := range slice {
			fmt.Println("Hello:", name)
		}
	}



	func salary(myFunc func(int) int){
		fmt.Println(myFunc(5))
	}


	calculateSalary := func(x int) int {
		return x *3
	}

	salary(calculateSalary)

	test := func(x int) int {
		return x * 2
	}(5)

	fmt.Println(test)



func printNames(arr [1]string) {
    fmt.Println("Hello:", arr[0])
}


func main() {

	var names [1]string


    fmt.Print("Skriv in ditt namn: ")

	fmt.Scanln(&names[0])

    printNames(names)
}


	/*
	// "Vanlig" for loop
	sum := 0
	for i := 1; i < 5; i++ {
 		sum += i
	}
	fmt.Println(sum) // 10 (1+2+3+4)




	// ”Infinite”
	infinateSum := 0
	for {
 		infinateSum++
 		if infinateSum > 100 {
			fmt.Println(infinateSum)
			break
 		}
	}


		n := 1
		for n < 5 {
 		n *= 2
		}
		fmt.Println(n) // 8 (1*2*2*2)







	strings := []string{"hello", "world"}
		for i, s := range strings {
 			fmt.Println(i, s)
		}


	---------------



	arr := []interface{}{1, "hej", 3.14, true}

	fmt.Println(arr)


	names := []string{}

    var name string


    fmt.Print("Skriv in ditt namn: ")


	fmt.Scanln(&name)

    names = append(names, name)

    fmt.Println("Hello:", names[0])




	var names [1] string

	fmt.Println("Enter your name: ")
	fmt.Scanln(&names[0])
	fmt.Println("Hello", names[0])




var name string;


	fmt.Println("Enter your name: ")
	fmt.Scanln(name)
	fmt.Println("Hello", name)




x := writeOutHello

	x()



/* Day 1 - Introduktion till Go

// Deklarera med var utanför funktioner, alternativ med rätt typ
var name = "Yahya"
var age = 34

// Med typer
var username string = "yahya"
var userAge int8 = 34;

	// I funktioner

	firstname := "Yahya"
	randomNumber := 3

	// Med formatsträng
	fmt.Printf("Hello! %s, your number is %d\n", firstname, randomNumber)

	// Eller enklare med Println
	fmt.Println("Hello!", firstname, "your number is", randomNumber)


	fmt.Println(username, userAge)
	fmt.Println(name, age)

*/