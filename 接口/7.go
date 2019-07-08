package main

import "fmt"

type interfaceonw interface {
	freekick()
}

type Bbb struct{}

func (bbb Bbb) freekick() {
	fmt.Println("调用接口一")
}

type Aaa struct{}

func (aaa Aaa) freekick() {
	fmt.Println("调用接口二")
}

func main() {
	var interfaceonw interfaceonw
	interfaceonw = new(Aaa)
	interfaceonw.freekick()

	interfaceonw = new(Bbb)
	interfaceonw.freekick()
}
