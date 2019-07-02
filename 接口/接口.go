package main

import "fmt"

/* 定义一个Phone接口 */
type Phone interface {
	call()
}

/* 结构体 */
type NokiaPhone struct {
}

/* 方法 */
func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

/* 结构体 */
type IPhone struct{}

/* 方法 */
func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main() {

	var phone Phone /* 接口名称 */

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

}
