package main

import (
	"fmt"
	"log"
	"net/http"
)

func HomeEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world :)")
}

func main() {
	http.HandleFunc("/", HomeEndpoint)
	fmt.Println("Server is running on port 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
