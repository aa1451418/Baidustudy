package main

import "fmt"

func main() {

	//切片
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(len(arr))
	fmt.Println(arr)

	//数组
	add := [2]int{1, 2}
	fmt.Println(cap(add))

	//make 定义切片
	acc := make([]int, 10)
	fmt.Println(len(acc))

}
