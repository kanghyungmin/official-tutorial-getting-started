package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

type Vertex struct {
	Lat, Long float64
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	return m
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	//Function closure
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)

	}

	//mutating maps
	// m := make(map[string]int)
	// m["Answer"] = 42
	// m["Answer"] = 48
	// delete(m, "Answer")
	// v, ok := m["Answer"]
	// fmt.Println("The value:", v, "Present?", ok)

	//Map literals
	// var m = map[string]Vertex{
	// 	"Bell Labs": Vertex{
	// 		40.68433, -74.39967,
	// 	},
	// 	"Google": Vertex{
	// 		37.42202, -122.08408,
	// 	},
	// }
	// fmt.Println(m)

	//Maps
	// var m map[string]Vertex

	// m = make(map[string]Vertex)
	// m["Bell Labs"] = Vertex{
	// 	40.68433, -74.39967,
	// }
	// fmt.Println(m["Bell Labs"])

	//Range continued
	// var pow []int = make([]int, 10)
	// for i := range pow {
	// 	pow[i] = 1 << uint(i)
	// }
	// for _, value := range pow {
	// 	fmt.Println(value)
	// }

	// var pow []int = []int{1, 2, 4, 8, 16, 32, 64, 128}
	// for i, v := range pow {
	// 	fmt.Printf("2**%d = %d\n", i, v)
	// }

	// var s []int
	// printSlice(s)

	// // append works on nil slices.
	// s = append(s, 0)
	// printSlice(s)

	// // The slice grows as needed.
	// s = append(s, 1)
	// printSlice(s)

	// // We can add more than one element at a time.
	// s = append(s, 2, 3, 4)
	// printSlice(s)
	// board := [][]string{
	// 	[]string{"_", "_", "_"},
	// 	[]string{"_", "_", "_"},
	// 	[]string{"_", "_", "_"},
	// }

	// board[0][0] = "X"
	// board[2][2] = "O"
	// board[1][2] = "X"
	// board[1][0] = "O"
	// board[0][2] = "X"

	// for i := 0; i < len(board); i++ {
	// 	fmt.Printf("%s\n", board[i])
	// }

	//Slice literals
	// q := []int{2, 3, 5, 7, 11, 13}
	// fmt.Println(q)

	// r := []bool{true, false, true, true, false, true}
	// fmt.Println(r)
	// s := []struct {
	// 	i int
	// 	b bool
	// }{
	// 	{2, true},
	// 	{3, false},
	// 	{5, true},
	// 	{7, true},
	// }
	// fmt.Println(s)

	// primes := [6]int{2, 3, 5, 7, 11, 13}
	// var s []int = primes[1:4]
	// fmt.Println(s)
	// var a [2]string
	// a[0] = "Hello"
	// a[1] = "World"
	// fmt.Println(a[0], a[1])
	// fmt.Println(a)

	// primes := [6]int{2, 3, 5, 7, 11, 13}
	// fmt.Println(primes)

	// var (
	// 	v1 = Vertex{1, 2}
	// 	v2 = Vertex{X: 1}
	// 	v3 = Vertex{}
	// 	p  = &Vertex{1, 2}
	// )
	// fmt.Println(v1, p, v2, v3)

	// v := Vertex{1, 2}
	// p := &v
	// p.X = 1e9
	// fmt.Println(v)
	// // i, j := 42, 2701

	// p := &i
	// fmt.Println(*p) //42
	// *p = 21
	// fmt.Println(i) //21

	// p = &j
	// *p = *p / 37   //2701/37
	// fmt.Println(j) //73

	// defer fmt.Println("world")
	// fmt.Println("hello")

	// fmt.Print("Go runs on.")

	// switch os := "darwin"; os {
	// case "darwin":
	// 	fmt.Println("OS X.", os)
	// case "linux":
	// 	fmt.Println("Linux.")

	// }

	// fmt.Println(sqrt(2), sqrt(-4))
	// fmt.Println(
	// 	pow(3, 2, 10),
	// 	pow(3, 3, 20),
	// )
	// sum := 0

	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)

	// sum := 1
	// for sum < 1000 {
	// 	sum += sum
	// }
	// fmt.Println(sum)
}
