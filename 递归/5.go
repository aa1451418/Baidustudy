package main

import "fmt"

func qee(z int) int {
	for z < 60 {
		return 3
		fmt.Println(z)
	}
	return z * qee(z*90)
}

func main() {
	var i int = 20
	fmt.Println(i, qee(i))

}
