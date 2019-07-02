package main

import "fmt"

func hsb(a, b, c, d int) int {
	return a + b - c/d
}

func main() {

	a := hsb(12, 45, 56, 77)
	fmt.Println("a=", a)

}
