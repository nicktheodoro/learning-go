package main

import (
	"errors"
	"fmt"
)

func main() {
	//INTEGER NUMBERS
	var num int = 100
	fmt.Println(num)

	var num2 uint = 100
	fmt.Println(num2)

	// alias

	// int32 = rune
	var num3 rune = 123456
	fmt.Println(num3)

	// uint8 = byte
	var num4 byte = 123
	fmt.Println(num4)

	//REAL NUMBERS
	var realNum1 float32 = 123.45
	fmt.Println(realNum1)

	var realNum2 float64 = 12300000000.45
	fmt.Println(realNum2)

	realNum3 := 12345.67
	fmt.Println(realNum3)

	//STRINGS
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	//BOLEAN
	var boolean1 bool = true
	fmt.Println(boolean1)

	var boolean2 bool
	fmt.Println(boolean2)

	//ERROR TYPE
	var err error = errors.New("Returning a new error")
	fmt.Println(err)
}
