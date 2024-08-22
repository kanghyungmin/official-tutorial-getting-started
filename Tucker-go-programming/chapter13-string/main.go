package main

import (
	"fmt"
	"strings"
	"unsafe"
)

type StringHeader struct {
	Data uintptr
	Len  int
}

func ToUpper1(str string) string {
	var rst string
	for _, s := range str {
		if s >= 'a' && s <= 'z' {
			rst += string('A' + (s - 'a'))
		} else {
			rst += string(s)
		}
	}
	return rst
}

func ToUpper2(str string) string {
	var builder strings.Builder
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			builder.WriteRune('A' + (c - 'a'))
		} else {
			builder.WriteRune(c)
		}
	}
	return builder.String()
}

func main() {
	// poet1 := "죽는 날까지 하늘을 우러러\n한점 부끄럼이 없기를,\n잎새에 이는 바람에도\n나는 괴로워했다.\n"
	// poet2 := `죽는 날까지 하늘을 우러러
	// 			한점 부끄럼이 없기를,
	// 			잎새에 이는 바람에도
	// 			나는 괴로워했다.
	// 			`
	// fmt.Println(poet1)
	// fmt.Println(poet2)

	// str := "Hello"

	// runes := []rune{72, 101, 108}

	// fmt.Println(len(str))
	// fmt.Println(len(runes))
	// fmt.Println(str)
	// fmt.Println(string(runes))
	// fmt.Println(unsafe.Sizeof(str))
	// fmt.Println(unsafe.Sizeof(runes))

	// str := "Hello, world!"
	// arr := []rune(str)

	// for i := 0; i < len(arr); i++ {
	// 	fmt.Printf("타입 : %T, 값 : %d, 문자 : %c\n", arr[i], arr[i], arr[i])
	// }

	var str string = "Hello, world!"

	var str1 string = ToUpper1(str)
	var str2 string = ToUpper1(str)

	fmt.Println(str1)
	fmt.Println(str2)

	stringHeader3 := (*StringHeader)(unsafe.Pointer(&str))
	stringHeader1 := (*StringHeader)(unsafe.Pointer(&str1))
	stringHeader2 := (*StringHeader)(unsafe.Pointer(&str2))

	fmt.Println(stringHeader1)
	fmt.Println(stringHeader2)
	fmt.Println(stringHeader3)
}
