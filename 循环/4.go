package main

import "fmt"

func main() {
	for i := 100; i >= 50; i++ {
		i = i + 2
		fmt.Println(i)
	}
}
