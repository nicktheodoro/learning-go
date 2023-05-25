package main

import "fmt"

func main() {
	// Declaring variable.
	var var1 string = "var1"
	fmt.Println(var1)

	// Declaring variable using type inference
	var2 := "var2"
	fmt.Println(var2)

	// Declaring multiples variables.
	var (
		var3 string = "var3"
		var4 string = "var4"
	)

	fmt.Println(var3, var4)

	// Declaring multiples variables using type inference.
	var5, var6 := "var5", "var6"
	fmt.Println(var5, var6)

	// Changing the variables values.
	var5, var6 = var6, var5
	fmt.Println(var5, var6)

	// Declaring constant.
	const const1 string = "const 1"
	fmt.Println(const1)
}
