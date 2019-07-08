package main

import "fmt"

func main() {

	/* 字符串类型数组 */
	a := []string{"我是字符串类型数组"}

	/* 整数型数组 */
	b := [4]int{1, 2, 4, 5}

	/* make函数数组 */
	c := make([]int, 20)
	d := make([]string, 5)
	d[1] = "我是make函数字符串类型"

	fmt.Println("a=", a, "\nb=", b, "\nc=", c, "\nd=", d)
}
