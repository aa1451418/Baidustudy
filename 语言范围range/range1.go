package main

import "fmt"

func main() {
	//整数型 不需要加 map
	onw := []int{0, 1, 2, 3, 4}
	for i := range onw {
		fmt.Println(onw[i])
	}

	//字体串类型必须加  map
	carcat := map[string]string{"car": "Cat", "Dog": "Banana", "Maney": "Onwmoney"}
	for car := range carcat {
		fmt.Println(carcat[car])
	}
}
