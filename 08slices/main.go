package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	var fruitList = []string{"Apple", "Mango"}

	fmt.Printf("Type of FruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Banana", "Orange")
	fmt.Println("New Fruit list is ", fruitList)

	fruitList = append(fruitList[1:3]) // last range is not included
	fmt.Println("New updated fruit list is ", fruitList)

	// alternate way to make a slice
	scores := make([]int, 4)

	scores[0] = 12
	scores[1] = 18
	scores[2] = 20
	scores[3] = 6

	// scores[4] = 19 will give error

	fmt.Println("Score slice -> ", scores)

	scores = append(scores, 14, 17)

	fmt.Println("New Score slice -> ", scores) // but this  will work

	// functions available with slices only
	sort.Ints(scores)
	fmt.Println("Sorted slice -> ", scores)
	fmt.Println("Is slice sorted -> ", sort.IntsAreSorted(scores))

	// remove value from slice

	var courses = []string{"react", "node", "go", "python", "c++", "java"}

	var index int = 2

	fmt.Println("Before removing index -> ", courses)

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("After removing index -> ", courses)
}