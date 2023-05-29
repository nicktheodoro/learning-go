package main

import "fmt"

func main() {
	num := 20
	fmt.Println(num)
	invertedNum := invertSign(num)
	fmt.Println(invertedNum)

	newNumber := 40
	fmt.Println(newNumber)
	invertSignWithPointer(&newNumber)
	fmt.Println(newNumber)
}

func invertSign(num int) int {
	return num * -1
}

func invertSignWithPointer(num *int) {
	*num = *num * -1
}
