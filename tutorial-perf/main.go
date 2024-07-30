package main

import (
	"fmt"
	"sync"
	"time"
)

type SageCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *SageCounter) Inc(key string) {
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
}

func (c *SageCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SageCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}

// import "fmt"
// func main() {
// 	ch := make(chan int, 2)
// 	ch <- 1
// 	ch <- 2
// 	ch <- 3
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// }

// import "fmt"

// func sum(s []int, c chan int) {
// 	sum := 0
// 	for _, v := range s {
// 		sum += v
// 	}
// 	c <- sum
// }

// func main() {
// 	s := []int{7, 2, 8, -9, 4, 0}
// 	c := make(chan int)
// 	go sum(s[:len(s)/2], c)
// 	go sum(s[len(s)/2:], c)
// 	x, y := <-c, <-c

// 	fmt.Println(x, y, x+y)
// }

// func say(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(100 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

// func main() {
// 	go say("world")
// 	say("hello")
// }
