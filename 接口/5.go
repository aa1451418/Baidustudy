package main

import "fmt"

type Iter interface {
	call()
}

type Hello struct{}

func (hello Hello) call() {
	fmt.Println("my is interface onw")
}

type World struct{}

func (world World) call() {
	fmt.Println("my is interface tow")

}

func main() {

	var Iter Iter
	Iter = new(Hello)
	Iter.call()

	Iter = new(World)
	Iter.call()

}
