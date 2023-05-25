package main

import "fmt"

func main() {
	person := map[string]string{
		"name":     "Nicolas",
		"lastName": "Theodoro",
	}

	fmt.Println(person)
	fmt.Println(person["name"])     // acessing key name
	fmt.Println(person["lastName"]) // acessing key lastName

	// Map nested structure
	person2 := map[string]map[string]string{
		"name": {
			"first": "Nicolas",
			"last":  "Theodoro",
		},
		"cursing": {
			"name":     "GO lang fundamentals",
			"duration": "170 hrs",
		},
	}

	fmt.Println(person2["cursing"]["name"]) // acessing nested key name
	delete(person2, "name")                 // deleting a key
	fmt.Println(person2)

	// adding new key
	person2["zodiacSign"] = map[string]string{
		"sun":       "Libra",
		"ascendant": "Aries",
	}
	fmt.Println(person2)
}
