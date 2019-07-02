package main

import "fmt"

func main() {
	var a map[string]string
	a = make(map[string]string) //必须要使用make函数来创建映射
	a["fa"] = "法国"

	for b := range a {
		fmt.Println(a[b])
	}

}
