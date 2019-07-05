package main

import "fmt"

func tre(b int) int {
	for b > 30 { //当b的值大于30侧返回20
		return 20
	}
	return b * tre(b+2)
}

func main() {
	var c = 39
	fmt.Println(c, tre(c))

}
