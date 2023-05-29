package main

import "fmt"

func main() {
	defer func1() // defer the execution until the last moment
	func2()
	fmt.Println(func3(5, 5))
}

func func1() {
	fmt.Println("Executing function 1")
}

func func2() {
	fmt.Println("Executing function 2")
}

func func3(n1, n2 float32) bool {
	defer fmt.Println("Average was calculated, the result will be returned")
	fmt.Println("Entering in the function")

	average := (n1 + n2) / 2

	return average >= 7
}
