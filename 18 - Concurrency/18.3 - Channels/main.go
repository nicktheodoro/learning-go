package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	go write("Hello World", channel)

	fmt.Println("After function wirite starts executing")

	for msg := range channel {
		fmt.Println(msg)
	}

	fmt.Println("End of program")
}

func write(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- text
		time.Sleep(time.Second)
	}

	close(channel)
}
