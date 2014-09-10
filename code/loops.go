package main

import (
	"fmt"
)

func main() {
	// START SLICE OMIT
	// Slice
	var names []string
	names = append(names, "Thomas")
	names = append(names, "Krampl")

	// Array
	sports := [3]string{"American Football", "Tennis"}
	sports[2] = "Handball"
	// END SLICE OMIT

	// START OMIT
	for {
		break
	}

	for 1 == 1 {
		break
	}

	for i := 0; i < len(sports); i++ {
		fmt.Println(i, sports[i])
	}

	for index, value := range names {
		fmt.Println(index, value)
	}
	// END OMIT
}
