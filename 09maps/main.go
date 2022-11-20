package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("Short forms of languages -> ", languages)
	fmt.Println("Short for Python -> ", languages["PY"])

	delete(languages, "RB")
	fmt.Println("Short forms after deletion -> ", languages)

	for key, value := range languages {
		fmt.Printf("For key %v value is %v\n", key, value)
	}
}