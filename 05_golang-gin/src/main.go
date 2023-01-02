package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Printf("Starting server at port 8000\n")

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Hello World!!")
	})
	http.ListenAndServe(":8000", nil)
}
