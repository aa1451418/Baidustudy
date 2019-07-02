package main

import "fmt"

func main() {

	//用make函数定义的切片
	a := make([]int, 2)
	b := make([]int, 2, 10)
	fmt.Println(a, b)
	fmt.Println(len(a))
	fmt.Println(cap(b))

	//数组
	arr := [3]int{1, 24}
	fmt.Println(arr)

	//切片
	sliceA := []int{1, 2, 3}
	fmt.Println(cap(sliceA))
}
