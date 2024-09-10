package main

import (
	"fmt"
	"hash/fnv"
	"maps"
	"strconv"

	"golang.org/x/exp/constraints"
)

func add[T constraints.Integer | constraints.Float](a, b T) T {
	return a + b
}

func Print[T any](a, b T) {
	fmt.Println(a, b)
}

func Print2(a, b interface{}) {
	fmt.Println(a, b)
}

// func main() {
// 	Print(1, 2)
// 	Print(3.14, 1.43)
// 	Print("hello", "world")
// 	Print("world", "hello")
// }

type Integer interface {
	~int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64
}

func add2[T Integer](a, b T) T {
	return a + b
}

type MyInt int

// func main() {
// 	add2(1, 2)
// 	var a MyInt = 3
// 	var b MyInt = 4
// 	add2(a, b)
// }

type ComparableHasher interface {
	comparable
	Hash() uint64
}

type MyString string

func (s MyString) Hash() uint64 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return uint64(h.Sum32())
}

func Equal[T ComparableHasher](a, b T) bool {
	return a.Hash() == b.Hash()
}

//	func main() {
//		var str1 MyString = "hello"
//		var str2 MyString = "Wolrd"
//		fmt.Println(Equal(str1, str2))
//	}
func Map[F, T any](s []F, f func(F) T) []T {
	rst := make([]T, len(s))
	for i, v := range s {
		rst[i] = f(v)
	}
	return rst
}

// func main() {
// 	doubled := Map([]int{1, 2, 3}, func(v int) int {
// 		return v * 2
// 	})

// 	uppered := Map([]string{"hello", "world"}, func(s string) string {
// 		return strings.ToUpper(s)
// 	})

// 	fmt.Println(doubled)
// 	fmt.Println(uppered)
// }

type Node[T any] struct {
	val  T
	next *Node[T]
}

func NewNode[T any](val T) *Node[T] {
	return &Node[T]{val: val}
}

func (n *Node[T]) Push(val T) *Node[T] {
	newNode := NewNode(val)
	n.next = newNode
	return newNode
}

// func main() {
// 	node1 := NewNode(1)
// 	node1.Push(2).Push(3).Push(4)

// 	for node1 != nil {
// 		fmt.Print(node1.val, " ")
// 		node1 = node1.next
// 	}
// 	fmt.Println()

// 	node2 := NewNode("Hi")
// 	node2.Push("Hello").Push("World")

// 	for node2 != nil {
// 		fmt.Print(node2.val, "->")
// 		node2 = node2.next
// 	}
// 	fmt.Println()
// }

type NodeType1 struct {
	val  interface{}
	next *NodeType1
}

type NodeType2[T any] struct {
	val  T
	next *NodeType2[T]
}

// func main() {
// 	node1 := &NodeType1{val: 1}
// 	node2 := &NodeType2[int]{val: 1}

// 	var v1 int = node1.val.(int)
// 	var v2 int = node2.val

// 	fmt.Println(v1, v2)
// }

// func main() {
// 	var v1 int = 3
// 	var v2 interface{} = v1
// 	fmt.Println(v1, v2)
// }

// func main() {
// 	names := []string{"Alice", "Bob", "Vera"}
// 	n, found := slices.BinarySearch(names, "Vera")

// 	fmt.Println("Vera is at index", n, "found:", found)

// 	n, found = slices.BinarySearch(names, "Bill")
// 	fmt.Println("Bill is at index", n, "found:", found)
// }

// func main() {
// 	type Person struct {
// 		Name string
// 		Age  int
// 	}
// 	people := []Person{
// 		{"Alice", 25},
// 		{"Bob", 30},
// 		{"Gopher", 13},
// 	}

// 	n, found := slices.BinarySearchFunc(people, "Bob", func(a Person, b string) int {
// 		return cmp.Compare(a.Name, b)
// 	})
// 	fmt.Println("Bob is at index", n, "found:", found)
// }

// func main() {
// 	m1 := map[string]int{
// 		"one": 1,
// 		"two": 2,
// 	}
// 	m2 := map[string]int{
// 		"one": 10,
// 	}

// 	maps.Copy(m2, m1)
// 	fmt.Println(m2)

// 	m2["one"] = 100
// 	fmt.Println("m1 is", m1)
// 	fmt.Println("m2 is", m2)

// 	m3 := map[string][]int{
// 		"one": {1, 2, 3},
// 		"tow": {4, 5, 6},
// 	}
// 	m4 := map[string][]int{
// 		"one": {10, 20, 30},
// 	}

// 	maps.Copy(m4, m3)
// 	fmt.Println(m4)

// }

func main() {

	m1 := map[int]string{
		1:   "1",
		10:  "10",
		100: "100",
	}
	m2 := map[int]int{
		1:   1,
		10:  10,
		100: 100,
	}

	eq := maps.EqualFunc(m1, m2, func(v1 string, v2 int) bool {
		i, err := strconv.Atoi(v1)
		if err != nil {
			// handle the error
			return false
		}
		// fmt.Println(i, v2)
		return v2 == i
	})

	fmt.Println(eq)
}
