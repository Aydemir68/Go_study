package main

import "fmt"

func ChangeString(str *string) {
	*str = "TEST"
}

func main() {
	s := "test"
	fmt.Println(s)
	//ChangeString(&s)
	fmt.Println(&s)

}
