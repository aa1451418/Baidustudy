package main

import "fmt"

/* 指针 */
func main() {

	var a = 10
	var b *int = &a
	fmt.Println(&b)

}
