package main

import "fmt"

func main() {

	msg := make(chan int)
	go func() {
		 msg <- 10 }()
	sgm := <-msg
	fmt.Println(sgm)

}
