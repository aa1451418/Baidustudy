package main

import "fmt"

func main() {

	a := []string{"我是切片"}
	fmt.Println(a)

	
	/* 数组 */
	c := [2]int{1, 2}
	fmt.Println(c)

	/* make函数切片 */
	b := make([]int, 10)
	fmt.Println(b)
}
