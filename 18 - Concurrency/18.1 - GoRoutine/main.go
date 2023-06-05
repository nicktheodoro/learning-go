package main

import (
	"fmt"
	"time"
)

func main() {
	go write("Hello World!") //goroutine
	write("Programming in Golang")
}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
