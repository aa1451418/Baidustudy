package main

import (
	"fmt"
	"time"
)

/* 方法 */
func say(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {

	go say("world")
	say("Hellow")

}
