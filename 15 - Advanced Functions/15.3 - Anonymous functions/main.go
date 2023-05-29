package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hello Wolrd")
	}()

	func(text string) {
		fmt.Println(text)
	}("Passing param")

	text := func(text string) string {
		return fmt.Sprintf("Receiving -> %s", text)
	}("Param")

	fmt.Println(text)
}
