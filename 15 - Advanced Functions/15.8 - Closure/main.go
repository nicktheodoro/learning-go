package main

import "fmt"

func main() {
	text := "Inside main function"
	fmt.Println(text)

	newFunction := closure()
	newFunction()
}

func closure() func() {
	text := "Inside closure function"

	function := func() {
		fmt.Println(text)
	}

	return function
}
