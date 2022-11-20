package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs")

	gurashish := User{"Gurashish Singh Gill", "gillgurashish171@gmail.com", 21, true}
	fmt.Println(gurashish)
	fmt.Printf("Name is %v and email is %v \n", gurashish.Name, gurashish.Email)
}

type User struct {
	Name string
	Email string
	Age int
	IsMale bool
}