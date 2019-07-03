package main

import "fmt"

const a int = 12

func main() {

	const s string = "我是常量"

	fmt.Println(s)
	fmt.Println(a)
	fmt.Println(float64(a))

}
