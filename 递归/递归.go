package main

import "fmt"

func flot(i int) int { //被15行调用
	if i < 1 {
		return 1
	}
	return i * flot(i-1)
}

func main() {

	var i int = 15
	fmt.Println(i, flot(i)) //flot(i)递归flot函数的数据

}
