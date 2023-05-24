package main

import (
	"fmt"
	"module/auxiliary"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Wiriting from main file")
	auxiliary.Write()

	erro := checkmail.ValidateFormat("123")
	fmt.Println(erro)
}
