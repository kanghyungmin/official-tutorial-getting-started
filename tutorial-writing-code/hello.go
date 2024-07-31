package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	c := make(chan int) // 채널 할당

	// 정렬을 고루틴에서 시작하고 완료되면 채널에 신호를 보냄
	go func() {
		list := []int{5, 2, 4, 6, 1, 3}
		sort.Ints(list)
		c <- 1 // 신호 전송; 값은 중요하지 않음
	}()

	doSomethingForAWhile()

	<-c
	fmt.Println("Sorting completed!")
}

func doSomethingForAWhile() {
	time.Sleep(2 * time.Second)
}
