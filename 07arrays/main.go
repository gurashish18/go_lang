package main

import "fmt"


func main() {
	fmt.Println("Welcome to arrays")

	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "banana"
	fruitList[3] = "orange"

	fmt.Println("Fruit list is ", fruitList)
	fmt.Println("Length of Fruit list is ", len(fruitList))

	var vegList = [5] string{"potato", "beas", "mushroom"}
	fmt.Println("veg list is ", vegList)
	fmt.Println("Length of veg list is ", len(vegList))
}