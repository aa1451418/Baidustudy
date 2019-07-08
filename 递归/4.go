package main

import "fmt"

func dg1(a int) int {
	for a < 10 {
		return a
	}
	return a * dg1(a/2)
}

func main() {

	var b int = 23
	fmt.Println(b, dg1(b))

}
