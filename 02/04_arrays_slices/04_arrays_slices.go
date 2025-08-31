package main

import (
	"fmt"
	"unsafe"
)

func IsLE() bool {
	var number int16 = 0x0001
	pointer := (*int8)(unsafe.Pointer(&number)) //unsafe.Pointer нужен, чтобы можно было переключиться между
	return *pointer == 1                        //int16 к int8
}

func IsBE() bool {
	return !IsLE()
}

func main() {
	s := [5]int{1, 2, 3, 4, 5}
	//s2:=s[0:2]
	s2 := s[0:2:2]

	s2 = append(s2, 444)
	fmt.Println("S1:", s, ", S2:", s2)

	if IsLE() {
		fmt.Println("Little Endian")
	} else {
		fmt.Println("Big Endian")
	}

	//Битовые свдиги
	var val int8 = 3
	val = val << 7
	fmt.Println(val)

	//Циклы
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	//For как While
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	//For with range
	for _, num := range s { // нижнее подчеркивание нужно, если не нужно ипользовать индекс. Если индекс нужен, просто пишем index
		fmt.Println(num)
	}

	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	for _, row := range matrix {
		for _, col := range row {
			fmt.Printf("%d", col)
		}
		fmt.Println()
	}
}
