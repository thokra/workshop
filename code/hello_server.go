package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {
	fmt.Println("Server listening on port :3030")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":3030", nil)
}
