package main

import "fmt"

func main() {

	var ys map[string]string
	ys = make(map[string]string) //必须要使用make函数来创建映射。

	ys["fg"] = "法国"

	//这是range范围
	for sy := range ys {
		fmt.Println(ys[sy])
	}
}
