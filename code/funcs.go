package main

import "fmt"

// START OMIT
func hello(first, last string) (string, error) {
	return fmt.Sprintln("Hello", first, last), nil
}

func bye(names ...string) {
	//
}

func main() {
	s, _ := hello("Ola", "Norman")
	fmt.Println(s)
}

// END OMIT
