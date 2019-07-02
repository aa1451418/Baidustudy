package main

import "fmt"

func main() {
	mmap := map[int]string{1: "i is dufalt map"}
	mmap[2] = "我同样是map"

	fmt.Println(mmap)

}
