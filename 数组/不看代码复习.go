package main

import "fmt"

func main() {
	//没有定义初始值
	var name [3]string
	name[0] = "yes0"
	name[1] = "yes1"
	name[2] = "yes2"
	fmt.Println(name)


	
	//初始化数组值直接在var行就开始赋值了
	var empo = [2]string{"我是初始化值", "我也是"}
	fmt.Println(empo)

}
