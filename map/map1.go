package main

import "fmt"

func main() {
	m := map[int]string{1: "我是map1", 2: "我也是2"} /* {这里的内容是初始值} */
	m[2] = "我也是map3"                            /* 这是后面紧接着的值上 */
	m[4] = "我也是map4"                            /* 这是后面紧接着的值上 */
	m[5] = "我也是map5"                            /* 这是后面紧接着的值上 */
	fmt.Printf("%+v\n", m)
	delete(m, 1)
}
