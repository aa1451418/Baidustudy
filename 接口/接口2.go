package main

import "fmt"

/* 接口名称 */
type Python interface {
	ty()
}

/* 结构体 */
type Studio struct{}

/* 实现方法 */
func (studio Studio) ty() {
	fmt.Println("我是接口一")
}

/* 结构体 */
type Student struct{}

/* 实现方法 */
func (student Student) ty() {
	fmt.Println("I am is interface tow")
}

/* 函数体 */
func main() {
	var Python Python
	Python = new(Studio)
	Python.ty()

	Python = new(Student)
	Python.ty()
}
