package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}

// type ErrNegativeSqrt float64

// func (e ErrNegativeSqrt) Error() string {
// 	return "cannot Sqrt negative number: " + fmt.Sprint(float64(e))
// }

// func Sqrt(x float64) (float64, error) {
// 	if x < 0 {
// 		return 0, ErrNegativeSqrt(x)
// 	}
// 	z := 1.0
// 	for i := 0; i < 10; i++ {
// 		z = z - (z*z-x)/(2*z)
// 	}
// 	return z, nil
// }

// type MyError struct {
// 	When time.Time
// 	What string
// }

// func (e *MyError) Error() string {
// 	return fmt.Sprintf("at %v, %s", e.When, e.What)
// }

// func run() error {
// 	return &MyError{
// 		time.Now(),
// 		"it didn't work",
// 	}
// }

// func main() {
// 	if err := run(); err != nil {
// 		fmt.Println(err)
// 	}
// }

//exercise-stringer
// package main
// import "fmt"
// type IPAddr [4]byte
// func (ipaddr IPAddr) String() string {
// 	return fmt.Sprintf("[%v.%v.%v.%v]", ipaddr[0], ipaddr[1], ipaddr[2], ipaddr[3])
// }

// func main() {
// 	hosts := map[string]IPAddr{
// 		"loopback":  {127, 0, 0, 1},
// 		"googleDNS": {8, 8, 8, 8},
// 	}
// 	for name, ip := range hosts {
// 		fmt.Printf("%v: %v\n", name, ip)
// 	}

// }

// package main

// import (
// 	"fmt"
// )

// type I interface {
// 	M()
// }

// type T struct {
// 	S string
// }

// func (t *T) M() {
// 	if t == nil {
// 		fmt.Println("<nil>")
// 		return
// 	}
// 	fmt.Println(t.S)
// }

// func main() {
// 	var i I

// 	var t *T
// 	i = t
// 	describe(i)
// 	i.M()

// 	i = &T{"hello"}
// 	describe(i)
// 	i.M()
// }

// func describe(i I) {
// 	fmt.Printf("(%v, %T)\n", i, i)
// }

// // type Vertex struct {
// // 	X, Y float64
// // }
// // type Abser interface {
// // 	Abs() float64
// // }

// // func (v *Vertex) Abs() float64 {
// // 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// // }

// // func AbsFunc(v Vertex) float64 {
// // 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// // }

// // func main() {
// // 	var a Abser
// // 	f := MyFloat(-math.Sqrt2)
// // 	v := Vertex{3, 4}

// // 	a = f
// // 	a = v
// // 	a = &v

// // 	fmt.Println(a.Abs())
// // }

// // type MyFloat float64

// // func (f MyFloat) Abs() float64 {
// // 	if f < 0 {
// // 		return float64(-f)
// // 	}
// // 	return float64(f)
// // }

// // func (v *Vertex) Scale(f float64) {
// // 	v.X = v.X * f
// // 	v.Y = v.Y * f
// // }

// // func ScaleFunc(v *Vertex, f float64) {
// // 	v.X = v.X * f
// // 	v.Y = v.Y * f
// // }

// // func main() {
// // 	v := Vertex{3, 4}
// // 	v.Scale(2)
// // 	ScaleFunc(&v, 10)

// // 	p := &Vertex{4, 3}
// // 	p.Scale(3)
// // 	ScaleFunc(p, 8)
// // }

// // type Vertex struct {
// // 	X, Y float64
// // }

// // func (v Vertex) Abs() float64 {
// // 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// // }

// // func (v Vertex) Scale(f float64) {
// // 	v.X = v.X * f
// // 	v.Y = v.Y * f
// // }

// // func main() {
// // 	v := Vertex{3, 4}
// // 	v.Scale(10)
// // 	fmt.Println(v.Abs())
// // }

// // type Vertex struct {
// // 	X, Y float64
// // }

// // func (v Vertex) Abs() float64 {
// // 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// // }

// // func main() {
// // 	v := Vertex{3, 4}
// // 	fmt.Println(v.Abs())
// // }
