package main

import "fmt"

func main() {

	y := map[int]string{1: "我是map"}
	y[2] = "我也是"
	fmt.Println(y)

}
