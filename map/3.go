package main

import "fmt"

func main() {
	a := map[int]string{1: "我是map"}
	a[1] = "map tow"

	fmt.Println(a)
}
