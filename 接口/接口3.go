package main

import "fmt"

type Ife interface {
	call()
}

type Over struct{}

func (over Over) call() {
	fmt.Println("我是接口作用")
}

func main() {
	var Ife Ife
	Ife = new(Over)
	Ife.call()

}
