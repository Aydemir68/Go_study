package main

import "fmt"

func main() {
	papers := map[string]int{
		"sber":     2000,
		"hh":       1000,
		"картошка": 10,
	}

	fmt.Println(papers)

	papers["x5"] = 200

	fmt.Println(papers)

	delete(papers, "картошка")

	fmt.Println(papers)

}
