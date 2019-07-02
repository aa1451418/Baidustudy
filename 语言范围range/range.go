package main

import "fmt"

func main() {

	number := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	for i := range number {
		fmt.Println(number[i])
	}

	coCap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}
	for l := range coCap {
		fmt.Println(coCap[l])
	}

}
