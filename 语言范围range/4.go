package main

import "fmt"

func main() {

	var a = []string{1: "美国", 2: "中国"}

	for b := range a {
		fmt.Println(a[b])
	}

	c := []int{1, 0, 2, 5, 3, 3, 4, 9, 5, 0, 6}
	for e := range c {
		fmt.Println(c[e])

	}

}
