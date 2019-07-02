package main

import "fmt"

func main() {

	//整数型Range范围
	a := []int{1, 2, 5, 7}
	for i := range a {
		fmt.Println(a[i])
	}

	//字符串类型map Range范围
	b := map[string]string{"cat1": "Cat1"}
	for e := range b { //定义一个变量回变量b
		fmt.Println(b[e]) //这里同时调用b,e
	}
}
