package main

import (
	"fmt"
)

func main() {
	// stdin := bufio.NewReader((os.Stdin))
	// for {
	// 	fmt.Println("Please input the number.")
	// 	var number int
	// 	_, err := fmt.Scanln(&number)

	// 	if err != nil {
	// 		fmt.Println("Please input the number(2).")
	// 		stdin.ReadString('\n')
	// 		continue
	// 	}

	// 	fmt.Printf("The number is %d\n", number)
	// 	if number%2 == 0 {
	// 		break
	// 	}
	// }
	for i := 0; i < 5; i++ {
		for j := 0; j < 5-i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
