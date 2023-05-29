package main

func main() {
	sum, subtraction := mathCalcs(10, 20)
	println(sum, subtraction)
}

func mathCalcs(n1, n2 int) (sum, subtraction int) {
	sum = n1 + n2
	subtraction = n1 - n2
	return
}
