package main

import "fmt"

func process (temp *int) {
	var value int = 200
	temp = &value

	fmt.Println(temp, *temp)
}


func main() {
	fmt.Println("Hello")

	var value int = 100
	var pointer *int = &value

	fmt.Println("Address: ", pointer)
	fmt.Println("value: ", *pointer)

	*pointer = 500

	fmt.Println(pointer, value, *pointer)

	process(pointer)

	fmt.Println(*pointer)

	//Массивы и срезы
	s:= [4]int{1,2,3,4}
	fmt.Println(s, "len = ", len(s), "capacity = ", cap(s))

	s2:= []int{1,2,3,4,5}
	fmt.Println(s2, "len = ", len(s2), "capacity = ", cap(s2))

	s3:=make([]int,5,500) //тип, длина, емкость
	fmt.Println(s3, "len = ", len(s3), "capacity = ", cap(s3))

	s4 := s[0:2] //срез для массива s2, при этом емкость остается как у s, т.к. s4 ссылается на s
	fmt.Println(s4, "len = ", len(s4), "capacity = ", cap(s4)) 
	s5 := s[0:2:2] // снова срез для массива s2, но третьим аргументом указываем емкость
	fmt.Println(s5, "len = ", len(s5), "capacity = ", cap(s5)) 
}