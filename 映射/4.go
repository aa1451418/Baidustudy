package main

import "fmt"

func main() {

	var a map[string]string
	a = make(map[string]string)
	a["america"] = "美国"
	fmt.Println(a)

	for b := range a {
		fmt.Println(a[b])

	}

}
