package main

import "fmt"

//这就是结构体
type student struct {
	Id    int
	Name  string
	Age   int
	Sex   string
	Email string
}

//实例化一个结构体
func main() {

	var a = new(student)

	a.Name = "小王"
	a.Id = 1
	a.Age = 21
	a.Sex = "你是女的"
	a.Email = "szkjtv@qq.com"
	fmt.Println(a)

}
