package main

import "fmt"

type Jgt interface {
	cls()
}

type Qwer struct{}

func (qwer Qwer) cls() {
	fmt.Println("interface onw ")
}

type Tgg struct{}

func (tgg Tgg) cls() {
	fmt.Println("interface tow")
}

func main() {
	var Jgt Jgt
	Jgt = new(Tgg)
	Jgt.cls()

	Jgt = new(Qwer)
	Jgt.cls()

}
