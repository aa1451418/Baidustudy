package main

import "fmt"

func main() {
	var a map[string]string
	a = make(map[string]string)
	a["hg"] = "韩国"

	for b := range a {
		fmt.Println(a[b])
	}

}
