package main

import (
	"fmt"
)

func main() {
	channel := make(chan string, 2)

	channel <- "Hello World"
	channel <- "Programming in Golang"

	msg := <-channel
	msg2 := <-channel

	fmt.Println(msg)
	fmt.Println(msg2)
}
