package main

import "fmt"

func main() {
	p := person{"Nicolas", "Theodoro", 29, 1.76}
	fmt.Println(p)

	s := student{p, "Golang Fundamentals"}
	fmt.Println(s)
}

type person struct {
	name     string
	lastName string
	age      int
	height   float32
}

type student struct {
	person
	course string
}
