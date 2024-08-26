package main

import (
	"fmt"
	"strings"
	"time"
)

type myString string

func (m myString) ToLower() myString {
	return myString(strings.ToLower(string(m)))
}

func (m myString) ToUpper() myString {
	return myString(strings.ToUpper(string(m)))
}

type Courier struct {
	Name string
}

type Product struct {
	Name  string
	Price int
	Id    int
}

type Parcel struct {
	Pdt           *Product
	ShippedTime   time.Time
	DeliveredTime time.Time
}

func (c *Courier) SendProduct(p *Product) *Parcel {
	return &Parcel{Pdt: p, ShippedTime: time.Now()}

}

func main() {
	msg := myString("Hello, World!")

	msg2 := msg.ToLower()
	msg3 := msg.ToUpper()

	fmt.Println(msg)
	fmt.Println(msg2)
	fmt.Println(msg3)
}
