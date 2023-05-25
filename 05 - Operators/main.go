package main

import "fmt"

func main() {
	// ARITMETICS
	sum := 1 + 2
	subtraction := 1 - 2
	division := 1 / 2
	multiplication := 1 * 2
	module := 1 % 2

	fmt.Println(sum, subtraction, division, multiplication, module)

	// ATRIBUITION
	var variable string = "String"
	variable2 := "String"
	fmt.Println(variable, variable2)

	// RELATIONAL
	fmt.Println("\n--RELATIONAL--")
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// LOGIC
	fmt.Println("\n--LOGIC--")
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(true != false)

	// UNARY
	fmt.Println("\n--UNARY--")
	num := 10
	num++
	fmt.Println(num)
	num--
	fmt.Println(num)
	num += 15
	fmt.Println(num)
	num -= 15
	fmt.Println(num)
	num *= 3
	fmt.Println(num)
	num /= 3
	fmt.Println(num)
	num %= 3
	fmt.Println(num)
}
