package main

import (
	"fmt"
)

func main() {
	var u user
	u.name = "Nicolas"
	u.age = 29
	fmt.Println(u)

	u2 := user{"Nicolas", 29, address{"Francisco Sá", 131}}
	fmt.Println(u2)

	u3 := user{name: "Nicolas"}
	fmt.Println(u3)
}

type user struct {
	name    string
	age     int
	address address
}

type address struct {
	street string
	number int
}
