package main

import (
	"fmt"
	"os"
)

// func sum(nums ...int) int {
// 	sum := 0

// 	fmt.Printf("nums 타입 : %T", nums)
// 	for _, v := range nums {
// 		sum += v
// 	}
// 	return sum
// }

// func main() {
// 	fmt.Println(sum(1, 2, 3, 4))
// 	fmt.Println(sum(1, 2))
// 	fmt.Println(sum())
// }

// func add(a, b int) int {
// 	return a + b
// }
// func mul(a, b int) int {
// 	return a * b
// }

// func getOperator(op string) func(int, int) int {
// 	if op == "+" {
// 		return add
// 	} else if op == "*" {
// 		return mul
// 	} else {
// 		return nil
// 	}
// }

// type opFunc func(int, int) int

// func main() {
// 	var oper opFunc
// 	oper = getOperator("*")

// 	var result = oper(3, 4)
// 	fmt.Println(result)
// }

type Writer func(string)

func writeHello(writer Writer) {
	writer("heel Wordld")
}

func main() {
	f, err := os.Create("text.txt")

	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}
	defer f.Close()

	writeHello(func(msg string) {
		fmt.Fscanln(f, msg)
	})

}
