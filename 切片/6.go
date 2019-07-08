package main

import "fmt"

func main() {

	/* 整数型数组 */
	a := [2]int{1, 23}
	fmt.Println(a)

	/* 字符串类型数组 */
	b := []string{"字符串类型数组"}
	fmt.Println(b)

	/* make函数型数组 */
	c := make([]int, 10)
	c[1] = 12
	c[2] = 13
	c[3] = 14
	c[4] = 15
	c[5] = 16
	fmt.Println(c)
}
