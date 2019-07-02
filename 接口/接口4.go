package main

import "fmt"

/* 接口 */
type Ministry interface {
	Thecher()
}

/* 结构体 */
type Showther struct{}

/* 方法 */
func (showther Showther) Thecher() {
	fmt.Println("我是教师")
}

/* 结构体 */
type Enth struct{}

/* 方法 */
func (eth Enth) Thecher() {
	fmt.Println("我也是教师")
}

/* main函数体 */
func main() {
	var Ministry Ministry
	Ministry = new(Showther)
	Ministry.Thecher()

	Ministry = new(Enth)
	Ministry.Thecher()

}
