package main

import "fmt"

func main() {

	a := 1

	switch a {
	case 1:
		fmt.Println("is tow")
	case 2:
		fmt.Println("is onw")
	case 3:
		fmt.Println("is thress")
	}

	b := 100

	switch b {
	case 1:
		fmt.Println("这个条件不是等于100错")
	case 2:
		fmt.Println("我是等于100")
	}

}
