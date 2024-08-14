package main

import "fmt"

const epsilon = 0.000001

func equal(a, b float64) bool {
	if a > b {
		if a-b <= epsilon {
			return true
		} else {
			return false
		}
	} else {
		if b-a <= epsilon {
			return true
		} else {
			return false
		}
	}
}

func main1() {
	var a float64 = 0.1
	b := 0.2
	c := 0.3

	fmt.Printf("%0.18f + %0.18f == %0.18f", a, b, c)
	fmt.Printf("%0.18f == %0.18f : %v", c, a+b, equal(a+b, c))
}
