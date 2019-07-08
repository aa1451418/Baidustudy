package main

import "fmt"

type kkkk interface {
	cccc()
}

type Lll struct{}

func (lll Lll) cccc() {
	fmt.Println("this'is interface onw")
}

type Cll struct{}

func (cll Cll) cccc() {
	fmt.Println("this'is interface tow")
}

func main() {

	var kkkk kkkk
	kkkk = new(Lll)
	kkkk.cccc()

	kkkk = new(Cll)
	kkkk.cccc()

}
