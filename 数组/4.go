package main

import "fmt"

func main() {
	var a = [12]string{1: "1", 2: "tow"}
	a[1] = "onw"
	a[2] = "tow"
	fmt.Println(a)

	/* 定义初始值的必须要用=号，如没初始值侧不需要=也可 */
	var b = [4]int{1, 2, 3, 4}
	b[1] = 1

	fmt.Println(b)
}
