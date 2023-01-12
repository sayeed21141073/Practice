package main

import (
	"fmt"
	"net/http"
)

func main() {
	// http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {

	// 	fmt.Fprint(w, "Hello World!")
	// })
	http.HandleFunc("/greet", greet)
	http.ListenAndServe(":8080", nil)
}

func greet(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Hello World!")
}
