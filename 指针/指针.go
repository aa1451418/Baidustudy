package main

import "fmt"

func main() {

	var a string = "工是指针"
	var b int = 5
	var pa *string = &a
	var pb *int = &b
	*pa = "我是接收指针"
	*pb = 4
	fmt.Println(pa)
	fmt.Printf("%+v\n", &pb)
}
