package main

import "fmt"

func main() {
	//定义没有初始化的数组字符串类型
	var student [5]string
	student[1] = "老王数组一"
	student[2] = "老王数组一"
	student[3] = "老王数组一"
	student[4] = "老王数组一"

	var course [4]int
	course[1] = 1
	course[2] = 2
	course[3] = 3

	var Only = [4]string{"a", "b"} //{"a", "b"}初始化数据
	Only[2] = "c"
	Only[3] = "d"
	fmt.Println(Only)

	fmt.Println(student)
	fmt.Println(course)

}
