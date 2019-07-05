package main

import "fmt"

func main() {

	a := []string{"我是默认字符串类型数组"}
	fmt.Println(a)

	b := [5]int{1, 2, 3, 4}
	fmt.Println("整数型数组", b)

	c := make([]int, 10)
	fmt.Println("make函数数组：", c)

}
