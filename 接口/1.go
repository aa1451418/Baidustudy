package main

import "fmt"

/* 接口 */
type jk interface {
	jjk()
}

type Jkk struct{}

func (jkk Jkk) jjk() {
	fmt.Println("我是接口1")
}

type Phone2 struct{}

func (phone2 Phone2) jjk() {
	fmt.Println("我是接口2")
}

func main() {
	var jk jk
	jk = new(Jkk)
	jk.jjk()

	jk = new(Phone2)
	jk.jjk()
}
