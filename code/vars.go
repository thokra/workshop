package main

import "fmt"

func main() {
	// START OMIT
	var address string

	var name string = "Thomas"
	var surname = "Krampl"
	occupation := "Go evangelist" // HL
	// END OMIT
	_ = address
	fmt.Println(name, surname, occupation)
}
