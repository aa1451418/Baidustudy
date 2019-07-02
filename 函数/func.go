package main

import "fmt"

/* 函数类型 */
func ts(a, b int) int {
	return a + b
}

/* 函数类型 */
func st(c, d int) int {
	return c * d
}

func main() {
	/* res变量指向ts函数 */
	res := ts(12, 45)
	fmt.Println("res=:", res)
	/* esr变量指向st函数 */
	esr := st(453564, 456548)
	fmt.Println("esr=:", esr)

}
