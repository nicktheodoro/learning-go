package main

import (
	"fmt"
)

func main() {
	result := sum(10, 20)
	fmt.Println(result)

	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("Function is a type in go")

	resultSum, resultSub := mathCalcs(10, 15)
	fmt.Println(resultSum, resultSub)

}

func sum(num1 int, num2 int) int {
	return num1 + num2
}

func mathCalcs(num1, num2 int) (int, int) {
	sum := num1 + num2
	sub := num1 - num2

	return sum, sub
}
