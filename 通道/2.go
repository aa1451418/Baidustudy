package main

import "fmt"

func main() {

	td := make(chan string)

	go func() { td <- "pingonw" }()

	sgm := <-td
	fmt.Println(sgm)

}
