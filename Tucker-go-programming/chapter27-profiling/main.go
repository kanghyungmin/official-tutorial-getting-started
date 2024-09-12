package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"

	"math/rand"
)

func Fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return Fib(n-1) + Fib(n-2)
	}
}

// func main() {
// 	f, err := os.Create("cpu.prof")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()

// 	pprof.StartCPUProfile(f)
// 	defer pprof.StopCPUProfile()
// 	fmt.Println(Fib(50))

// 	time.Sleep(10 * time.Second)
// }

func main() {
	http.HandleFunc("/log", logHandler)
	http.ListenAndServe(":8080", nil)

}

func logHandler(w http.ResponseWriter, r *http.Request) {
	ch := make(chan int)
	go func() {
		time.Sleep(time.Duration(rand.Intn(400)) * time.Millisecond)
		ch <- http.StatusOK
	}()

	select {
	case status := <-ch:
		w.WriteHeader(status)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("timeout")
		w.WriteHeader(http.StatusRequestTimeout)
	}

}
