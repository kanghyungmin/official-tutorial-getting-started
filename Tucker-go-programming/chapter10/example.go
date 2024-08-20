package main

import "fmt"

func main() {
	// var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 25.5}

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(t[i])
	// }
	nums := [...]int{10, 20, 30, 40, 50}
	nums[2] = 300

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 25.5}

	for i, v := range t {
		fmt.Println(i, v)
	}

	var a [2][5]int = [2][5]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}}

	for _, arr := range a {
		for _, v := range arr {
			fmt.Print(v, " ")
		}
	}

}
