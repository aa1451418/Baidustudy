package main

import "fmt"

func main() {

	//切片
	a := []string{"i is 切片"}
	fmt.Println(a)

	//数组
	b := [2]int{1, 2}
	fmt.Println(b)

	//make函数切片
	c := make([]int, 20)
	fmt.Println(c)

}
