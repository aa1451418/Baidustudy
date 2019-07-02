package main

import "fmt"

func main() {
	a := map[int]string{1: "我是map"}
	a[0] = "the is map "

	fmt.Println(a)

}
