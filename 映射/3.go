package main

import "fmt"

func main() {

	var Editor map[string]string
	Editor = make(map[string]string) //使用make函数来创建映射
	Editor["nal"] = "China"

	for i := range Editor {
		fmt.Println(Editor[i])
	}

}
