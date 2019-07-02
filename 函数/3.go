package main

import "fmt"

func look(a, c int) int {
	return a * c
}

func main() {
	abc := look(1345, 113543654)
	fmt.Println(abc)
}
