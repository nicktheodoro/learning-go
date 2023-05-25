package main

import (
	"fmt"
)

func main() {

	var arr1 [5]int
	arr1[0] = 1
	fmt.Println(arr1)

	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	arr3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr3)

}
