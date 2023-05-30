package main

import "fmt"

func main() {
	generic("string")
	generic(1)
	generic(true)

	m := map[interface{}]interface{}{
		1:            "String",
		float32(100): true,
		"String":     "String",
		true:         float64(12.1),
	}

	fmt.Println(m)
}

func generic(i interface{}) {
	fmt.Println(i)
}
