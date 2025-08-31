package main

import "fmt"

func mathOperations(a int, b int) (sum int, sub int, mul int, div float32) {
	sum = a + b
	sub = a - b
	mul = a * b
	div = float32(a) / float32(b)

	return
}

func main() {
	a, b, c, d := mathOperations(4, 5)
	fmt.Println(a, b, c, d)
}
