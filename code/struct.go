package main

import (
	"encoding/json"

	"fmt"
)

// START OMIT
type User struct {
	Age            int
	Username, Name string
	admin          bool
}

func main() {
	usr := User{
		Age:      24,
		Username: "thokra",
		Name:     "Thomas Krampl",
		admin:    true,
	}
	// Samme som User{24, "thokra", "Thomas Krampl", true}

	fmt.Println(usr.Age, usr.Username)
	fmt.Println(usr)
}

// END OMIT

// START FUNC OMIT
func (u *User) String() string {
	d, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(d)
}

// END FUNC OMIT
