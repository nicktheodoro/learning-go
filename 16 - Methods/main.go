package main

import "fmt"

func main() {
	user1 := user{"User 1", 20}
	fmt.Println(user1)
	user1.save()
	fmt.Println(user1.ofLegalAge())
	user1.makeBirthday()
	fmt.Println(user1.age)
}

type user struct {
	name string
	age  uint
}

func (u user) save() {
	fmt.Printf("Saving the user %s on database", u.name)
}

func (u user) ofLegalAge() bool {
	return u.age >= 18
}

func (u *user) makeBirthday() {
	u.age++
}
