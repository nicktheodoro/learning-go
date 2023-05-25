package main

import "fmt"

func main() {
	var var1 int = 10
	var var2 int = var1

	fmt.Println(var1, var2)

	var1++
	fmt.Println(var1, var2)

	//  A pointer holds the memory address of a value
	var var3 int = 100
	var pointer *int = &var3 // set i through the pointer p

	fmt.Println(*pointer) // read i through the pointer p

	//This is known as "dereferencing" or "indirecting"!!!
	var3++
	fmt.Println(var3, *pointer) // var3 and pointer has the same value by dereferencing
}
