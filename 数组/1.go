package main

import "fmt"

func main() {

	var a [12]int /* 没有初始化数组 */
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println(a)

	var b = [2]string{"数组1", "数组2"} /* 初始化数组 */
	fmt.Println(b)

}
