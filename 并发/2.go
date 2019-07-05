package main

import "fmt"

/* 这个练习是错乱的当熟练代码就好了 */
func hsdy(i int) int {
	for i := 0; i < 10; i++ {
		i := i + 1
		fmt.Println("我不是并发", i)
	}
	return i
}

func main() {
	// for i := 0; i < 50; i++ {
	// 	i := i + 1
	// 	fmt.Println(i)
	// }
	res := hsdy(500)
	go fmt.Println(res)

}
