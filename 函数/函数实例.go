package main

import "fmt"

func hs(a, b int) int {
	return a + b
}

func sh(a, b, c int) int {
	return a + b - c
}


func main() {
	res := hs(1, 3)
	fmt.Println("res=", res)

	res1 := sh(4, 5, 6)
	fmt.Println("res1=", res1)

}
