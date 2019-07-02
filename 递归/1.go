package main

func ther(a int) int {
	for a < 1 {
		return 1
	}
	return a * ther(a-1)
}

func main() {
	var i int = 12
	println(i, ther(i))

}
