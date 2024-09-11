package main

import (
	"errors"
	"fmt"
	"strings"
)

func square(x int) int {
	return x * x
}

// func main() {
// 	fmt.Printf("9*9=%d\n", square(9))
// }

func fibonacci1(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return n
	}
	return fibonacci1(n-1) + fibonacci1(n-2)
}

func fibonacci2(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return n
	}
	one := 1
	two := 0
	rst := 0

	for i := 2; i <= n; i++ {
		rst = one + two
		two = one
		one = rst
	}
	return rst
}

func Atoi(input string) (int, error) {
	rst := 0
	negative := false
	input = strings.TrimSpace(input)
	if input[0] == '-' {
		negative = true
		input = input[1:]
	}

	for _, c := range input {
		if c < '0' || c > '9' {
			return 0, errors.New("not a number")
		}
		rst = rst*10 + int(c-'0')
	}

	if negative {
		rst = -rst
	}

	return rst, nil
}

func main() {
	fmt.Println(Atoi("123"))
	// fmt.Println(fibonacci2(13))
}
