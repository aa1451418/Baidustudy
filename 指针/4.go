package main

import "fmt"

/* 指针 */
func main() {

	var a = 1
	var b *int = &a   
	fmt.Println(&b)

}
