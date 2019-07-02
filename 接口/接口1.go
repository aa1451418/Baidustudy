package main

import "fmt"

/* 接口 */
type Phone1 interface {
	cal()
}

/* 结构体首字母必须大写 */
type IPhone1 struct {
}

func (iphone1 IPhone1) cal() {
	fmt.Println("你可以打电话给我")
}

/* 结构体和方法一起 */
type Hone struct {
}

func (hone Hone) cal() {
	fmt.Println("我可以打电话给你吗")
}

func main() {

	var Phone1 Phone1     //定义接口变量
	Phone1 = new(IPhone1) //新建结构体
	Phone1.cal()          //输出内容

	Phone1 = new(Hone) //新建结构体
	Phone1.cal()       //输出内容

}
