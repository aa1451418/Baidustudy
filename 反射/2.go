package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello learm English </h1>")
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Start Serverce...")
	http.ListenAndServe(":8000", nil)

}
