package main

import "fmt"

/* 指针 */
func main() {

	var a int = 13
	var b *int = &a
	fmt.Println(&b)

}
