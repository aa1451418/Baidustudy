package main

import "fmt"

func main() {
	//字符串类型数组
	a := []string{"我是切片"}
	fmt.Println(a)

	/* 整数型数组 */
	c := [2]int{1, 2}
	fmt.Println(c)

	/* make函数切片 */
	b := make([]int, 10)
	fmt.Println(b)
}
