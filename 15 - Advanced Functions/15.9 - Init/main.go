package main

import "fmt"

func main() {
	fmt.Println("Executing function main")
}

func init() {
	fmt.Println("Executing function init") // Init function always runs before first
}
