package main

import "fmt"

func Divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

func F(n int) int {
	if n < 0 {
		return n
	} else {
		return F(n-1) * F(n-2)
	}
}

func main() {
	fmt.Println(F(9))
}
