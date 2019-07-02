package main

import "fmt"

func main() {

	/* 整数型数组 */
	var a [2]int //没有初始化数组
	a[0] = 1000  //定义数组内容
	a[1] = 1     //定义数组内容
	fmt.Println(a)

	/* 字符串类型数组 */
	var b = [4]string{"精妙绝伦", "Message"} //这是初始化内容{"精妙绝伦", "Message"}
	b[0] = "America"                     //定义数组内容
	b[1] = "China"                       //定义数组内容
	b[2] = "Canada"                      //定义数组内容
	b[3] = "Jpeg"                        //定义数组内容
	fmt.Println(b)
}
