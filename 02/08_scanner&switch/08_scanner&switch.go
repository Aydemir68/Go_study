package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/k0kubun/pp"
)

func main () {
	scanner:=bufio.NewScanner(os.Stdin)
	

	for {
		fmt.Println("Enter command: 1 - Add, 2 - Delete, 3 - Exit")
	
		if ok:=scanner.Scan(); !ok {
			fmt.Println("Error!")
			return
		}

		text:=scanner.Text()

		fields:=strings.Fields(text)
		pp.Println(fields)

		fmt.Println(fields[0])


		switch text {
		case "3":
			fmt.Println("goodluck!")
			return
		case "1":
			text = ""
			fmt.Println("What you want to add? ")
			scanner:=bufio.NewScanner(os.Stdin)
			if ok:=scanner.Scan(); !ok {
				fmt.Println("Error!")
				return
		}
			text:=scanner.Text()
			fmt.Println("You wanna add: ", text)
		case "2":
			text = ""
			fmt.Println("What you want to delete? ")
			scanner:=bufio.NewScanner(os.Stdin)
			if ok:=scanner.Scan(); !ok {
				fmt.Println("Error!")
				return
		}
			text:=scanner.Text()
			fmt.Println("You wanna delete: ", text)
		default:
			fmt.Println("Unknown command:", text)
		}

	}

}