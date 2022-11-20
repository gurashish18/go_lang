package main

import "fmt"

const LoginToken string = "sdvsdulbvsd" // public

func main() {
	var username string = "gurashish"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedin bool = false
	fmt.Println(isLoggedin)
	fmt.Printf("Variable is of type: %T \n", isLoggedin)

	var number int = 9745
	fmt.Println(number)
	fmt.Printf("Variable is of type: %T \n", number)

	// default values and aliases

	var anothernumber int
	fmt.Println(anothernumber)

	var anotherstring string
	fmt.Println(anotherstring)

	var anotherboolean bool
	fmt.Println(anotherboolean)

	var anotherfloat float32
	fmt.Println(anotherfloat)

	// implicit type

	var website = "gurashish.com"
	fmt.Println(website)

	// no var style (can be used only inside a method)
	numberOfusers := 100
	fmt.Println(numberOfusers)

	fmt.Println(LoginToken)
}