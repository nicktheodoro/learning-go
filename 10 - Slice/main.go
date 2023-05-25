package main

import (
	"fmt"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	slice := []int{10, 11, 12, 13, 14, 15}
	fmt.Println(slice)

	slice = append(slice, 99)
	fmt.Println(slice)

	// Works like pointer.
	slice2 := arr[1:3]
	fmt.Println(slice2)

	arr[1] = 100
	fmt.Println(slice2)

	// Internal arrays
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacity

	// Increasing the slice capacity
	slice3 = append(slice3, 55)
	slice3 = append(slice3, 66)

	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	slice4 = append(slice4, 77)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
