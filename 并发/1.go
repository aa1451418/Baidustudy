package main

import "fmt"

func bf(s int) int {
	for i := 0; i < 10; i++ {
		s = i + 1
		fmt.Println(i)
	}
	return s
}

func main() {
	for i := 0; i < 50; i++ {
		i = i + 1
		go fmt.Println(i)
	}
}
