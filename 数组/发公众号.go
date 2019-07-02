package main

import "fmt"

func main() {

	//定义一个没有初始化的数组
	var a [2]string
	a[0] = "我是数组一"
	a[1] = "我是数组二"
	fmt.Println(a)

	//定义一个有初始化值的数组
	var b = [4]int{0: 1, 1: 2}
	fmt.Println(b)
}
