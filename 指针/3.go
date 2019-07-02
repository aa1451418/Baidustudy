package main

import "fmt"

func main() {
	var a = 1234
	var c *int = &a
	fmt.Println(&c)

}
