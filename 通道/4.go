package main

import "fmt"

func main() {
	as := make(chan string)
	go func() {
		as <- "我是通道"
	}()

	sa := <-as
	fmt.Println(sa)
}
