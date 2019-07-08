package main

import "fmt"

func main() {
	var c map[string]string

	c = make(map[string]string)
	c["去英国"] = "amog"

	for z := range c {
		fmt.Println(c[z])
	}
}
