package main

import "fmt"

func main() {
	total := sum(1, 2, 3, 4, 5, 200, 300)
	fmt.Println(total)

	write("Hello World ", 1, 2, 3, 4, 5)
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	return total
}

func write(text string, nums ...int) {
	for _, num := range nums {
		fmt.Println(text, num)
	}
}
