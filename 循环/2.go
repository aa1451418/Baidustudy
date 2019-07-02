package main

import "fmt"

func main() {
	for i := 0; i < 50; i++ { //第1个i是赋值，第2个i是判断第一个i的条件，第三个i是执行条件操作
		i = i + 1
		fmt.Println(i)
	}
}
