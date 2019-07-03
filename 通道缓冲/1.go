package main

import "fmt"

func main() {

	maeesge := make(chan string, 2)
	maeesge <- "buffered"
	maeesge <- "channel"

	fmt.Println(<-maeesge)
	fmt.Println(<-maeesge)

}
