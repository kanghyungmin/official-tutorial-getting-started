package main

import (
	"fmt"
	"unsafe"
)

// type House struct {
// 	Address string
// 	Size    int
// 	Price   float64
// 	Type    string
// }

// func main() {
// 	var house House
// 	house.Address = "Seoul"
// 	house.Size = 28
// 	house.Price = 9.8
// 	house.Type = "Apt"

// 	fmt.Println("Address: ", house.Address)
// 	fmt.Println("Size: ", house.Size)
// 	fmt.Println("Price: ", house.Price)
// 	fmt.Println("Type: ", house.Type)

// }

// type User struct {
// 	Name string
// 	ID   string
// 	Age  int
// }

type VIPUser struct {
	User
	VIPLevel int
	Price    int
}

type User struct {
	A int8
	C int8
	E int8
	B int

	D int
}

func main() {
	// user := User{"Jang", "jang", 32}
	// var vip VIPUser = VIPUser{
	// 	User{"Kang", "kang", 23},
	// 	3,
	// 	250,
	// }

	// println(user.Name)
	// println(vip.Name)
	// println(vip.ID)
	// println(vip.Age)
	// println(vip.VIPLevel)
	// println(vip.Price)

	user := User{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(user))

}
