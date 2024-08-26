package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age  int
}
type Students []Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	// var slice = []int{1, 2, 3}

	// slice2 := append(slice, 4)

	// fmt.Println(slice)
	// fmt.Println(slice2)

	// fmt.Println(&slice[0], &slice2[0])

	// slice1 := make([]int, 3, 5)
	// slice2 := append(slice1, 4, 5)

	// fmt.Println("slice1:", slice1, len(slice1), cap(slice1), &slice1[0])
	// fmt.Println("slice2:", slice2, len(slice2), cap(slice2), &slice2[0])

	// slice1[1] = 100

	// fmt.Println("After change")
	// fmt.Println("slice1:", slice1, len(slice1), cap(slice1), &slice1[0])
	// fmt.Println("slice2:", slice2, len(slice2), cap(slice2), &slice2[0])

	// slice1 = append(slice1, 500)
	// fmt.Println("After append")
	// fmt.Println("slice1:", slice1, len(slice1), cap(slice1), &slice1[0])
	// fmt.Println("slice2:", slice2, len(slice2), cap(slice2), &slice2[0])

	// slice1 := []int{1, 2, 3}
	// slice2 := append(slice1, 4, 5)

	// fmt.Println("slice1:", slice1, len(slice1), cap(slice1), &slice1[0])
	// fmt.Println("slice2:", slice2, len(slice2), cap(slice2), &slice2[0])

	// array := [5]int{1, 2, 3, 4, 5}
	// slice := array[0:2]

	// fmt.Println("array:", array, &array[0])
	// fmt.Println("slice: ", slice, len(slice), cap(slice), &slice[0])

	// slice[0] = 100
	// fmt.Println("After change slice[0] = 100")
	// fmt.Println("array:", array, &array[0])
	// fmt.Println("slice: ", slice, len(slice), cap(slice), &slice[0])

	// slice = append(slice, 6, 7)
	// fmt.Println("After append")
	// fmt.Println("array:", array, &array[0])
	// fmt.Println("slice: ", slice, len(slice), cap(slice), &slice[0])

	// slice := []int{1, 2, 3, 4, 5, 6}
	// idx := 2

	// slice = append(slice[:idx], slice[idx+1:]...)

	// fmt.Println(slice)

	// s := []int{1, 2, 3, 4, 5}
	// sort.Ints(s)

	// fmt.Println(s)

	s := []Student{
		{"Maria", 20},
		{"Andrew", 22},
		{"John", 21},
	}
	sort.Sort(Students(s))
	fmt.Println(s)

}
