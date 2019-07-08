package main

import "fmt"

func main() {

	tdh := make(chan string, 2)
	tdh <- "缓存一组"
	tdh <- "缓存二组"
	fmt.Println(tdh)
	tdhz := make(chan int, 4)
	tdhz <- 1
	tdhz <- 2
	tdhz <- 3
	tdhz <- 4
	fmt.Println(tdhz)

}
