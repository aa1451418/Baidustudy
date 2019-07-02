package main

import "fmt"

func main() {
	a := 255
	var b *int = &a
	fmt.Println(b)

	//指针多种声明方式
	c := 255   //自动推导类型 :=
	f := 255   //自动推导类型 ：=
	d := c + f //自动推导类型 ：=
	g := &d    //自动推导类型 ：=
	fmt.Println(c)
	fmt.Println(f)
	fmt.Println(d)
	fmt.Println(*g)

	h := "我是指针咬我啊" //自动推导类型 ：=
	j := &h        //自动推导类型 ：=
	fmt.Println(*j)

}
