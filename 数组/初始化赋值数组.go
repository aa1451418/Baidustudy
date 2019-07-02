package main

import "fmt"

func main() {
	//给数组赋初始化值,a,b,c
	var student = [5]string{"a", "b", "c"}
	fmt.Println(student)

	//给数组赋初始化值,1,2,4,6,6
	var dapartment = [5]int{1, 2, 4, 6, 6}
	fmt.Println(dapartment)

	//这是没有初始化的数组
	var conurse [5]int
	conurse[0] = 10
	conurse[2] = 20
	conurse[3] = 30
	conurse[4] = 40
	fmt.Println(conurse)

	//这是没有初始化的数组
	var school [2]int
	school[0] = 100
	school[1] = 200
	fmt.Println(school)

}
