package main

import "fmt"

func main() {

	//整数型直接操作就好了
	o := []int{1, 3, 4, 6, 7}

	for i := range o {
		fmt.Println(o[i])
	}

	//字符串类型必须添加  map作为函数
	str := map[string]string{"dog": "dog", "office": "Office"}
	for b := range str {
		fmt.Println(str[b])
	}

}
