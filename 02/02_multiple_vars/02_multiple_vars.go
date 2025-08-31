package main

import "fmt"

var price2 = 2.5

func main() {

	var a = "Amal"

	var b = "ekum"

	a1 := 1
	a2 := 2
	price := 2.50

	fmt.Println("Test", a+b, a1+a2, price2)

	const name, age = "John", 24
	fmt.Printf("Hi! My name is %s, I'm %d yeard old. $%.2f\n", name, age, price)
}
