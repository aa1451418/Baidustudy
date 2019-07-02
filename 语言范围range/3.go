package main

import "fmt"

func main() {

	onw := []int{1, 2, 3}
	for a := range onw {
		fmt.Println(onw[a])
	}

	emp := map[string]string{"字符串1": "我就是字符串"}
	for b := range emp {
		fmt.Println(emp[b])
	}

}
