package main

import "fmt"

func main() {
	num := 10

	if num > 15 {
		fmt.Println("Number greater than 15")
	} else {
		fmt.Println("Number less than 15")
	}

	if num2 := num; num2 > 5 {
		fmt.Println("Number greater than 5")
	} else if num2 > 0 {
		fmt.Println("Number greater than 0")
	} else {
		fmt.Println("Number less than 0")
	}
}
