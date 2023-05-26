package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i <= 10 {
		fmt.Println(i)
		// time.Sleep(time.Second)
		i++
	}

	for j := 0; j <= 10; j++ {
		fmt.Println(j)
		// time.Sleep(time.Second)
	}

	lenguages := [3]string{"Go", "Node", "Java"}

	for index, value := range lenguages {
		println(index, value)
		// time.Sleep(time.Second)
	}

	for index, value := range "WORD" {
		println(index, string(value))
		// time.Sleep(time.Second)
	}

	user := map[string]string{
		"name":     "Nicolas",
		"lastName": "Theodoro",
	}

	for key, value := range user {
		println(key, value)
		time.Sleep(time.Second)
	}
}
