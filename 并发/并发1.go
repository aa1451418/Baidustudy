package main

import (
	"fmt"
	"time"
)

func tilk(t string) {

	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Millisecond)
		fmt.Println(t)
	}
}

func main() {
	go tilk("Hellow")
	tilk("world")
}
