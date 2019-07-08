package main

import "fmt"

func main() {
	b := 10

	switch b {
	case 1:
		fmt.Println("b=1")
	case 2:
		fmt.Println("b!=1")
	case 10:
		fmt.Println("b=10")
	}

}
