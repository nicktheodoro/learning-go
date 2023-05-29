package main

import "fmt"

func main() {
	fmt.Println(average(7, 7))
	fmt.Println("After execution")
}

func average(n1, n2 float32) bool {
	defer recoverExecution()
	average := (n1 + n2) / 2

	if average > 7 {
		return true
	}

	if average < 7 {
		return false
	}

	panic("The average is exactly 7") // If panic is called, program execution is stopped
}

func recoverExecution() {
	if r := recover(); r != nil { // If you use recover, program execution doesn't stop
		fmt.Println("Execution successfully recovered")
	}
}
