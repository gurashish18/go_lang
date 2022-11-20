package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")

	var number int = 22

	var ptr = &number

	fmt.Println("Value of pointer is ", ptr)
	fmt.Println("Value at pointer is ", *ptr)

	*ptr = *ptr + 5
	fmt.Println("New value is ", number)
}