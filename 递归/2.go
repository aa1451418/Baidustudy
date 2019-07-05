package main

import "fmt"

func qwe(a int) int {
	for a < 20 {
		return 1
		fmt.Println(a)
	}
	return a * qwe(a-1)
}

func main() {
	var i int = 20
	fmt.Println(i, qwe(i))

}
