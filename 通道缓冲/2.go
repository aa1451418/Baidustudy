package main

import "fmt"

func main() {

	td1 := make(chan string, 2) //chan通道缓存
	td1 <- "onw"
	td1 <- "tow"

	fmt.Println(td1)
	fmt.Println(td1)

}
