package main

func ther(a int) int {
	for a < 1 {
		return 1
	}
	return a * ther(a-1)
}

func main() {
	var i int = 12
	println(i, ther(i)) //调用第三3行ther，然后第7行也要用同样的函数回参

}
