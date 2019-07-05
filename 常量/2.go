package main

import "fmt"

const b int = 134

func main() {
	const a string = "我是常量"

	fmt.Println(a)
	fmt.Println("我是全局常量", b)

}
