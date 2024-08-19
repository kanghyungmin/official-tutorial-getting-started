package main

import "fmt"

func getMyage() (int, bool) {
	return 33, true
}

func main() {
	if age, ok := getMyage(); ok && age < 30 {
		fmt.Println("You are young!", age)
	} else if ok && age < 30 {
		fmt.Println("You are in middle age!", age)
	} else if ok {
		fmt.Println("You are beautiful!", age)
	} else {
		fmt.Println("You are old!", age)
	}
	//fmt.Println("Your age is ", age)
}
