package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(4)

	go func() {
		write("Hello World!")
		waitGroup.Done()
	}()

	go func() {
		write("Programming in Golang")
		waitGroup.Done()
	}()

	go func() {
		write("GoRoutine 3")
		waitGroup.Done()
	}()

	go func() {
		write("GoRoutine 4")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func write(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
