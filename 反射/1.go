package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>World News</h1>")
}

func main() {

	http.HandleFunc("/", index)
	fmt.Println("Service...")
	http.ListenAndServe(":4005", nil)

}
