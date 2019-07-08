package main

import "fmt"

func main() {

	a := []int{1, 24, 5, 6}
	for b := range a {
		fmt.Println(a[b])
	}

	c := []string{"法国", "中国", "尔罗斯", "美国"}
	for d := range c {
		fmt.Println(c[d])
	}

}
