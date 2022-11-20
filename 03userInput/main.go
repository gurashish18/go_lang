package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Hi There welcome!!!"
	fmt.Println(welcome)

	fmt.Println("Enter a review of our pizza")
	// comma ok syntax

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for giving review: ", input)
}