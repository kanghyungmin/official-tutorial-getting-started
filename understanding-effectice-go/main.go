package main

import (
	"log"
)

type Work struct {
	Task string
}

func server(workChan <-chan *Work) {
	for work := range workChan {
		go safelyDo(work)
	}
}

func safelyDo(work *Work) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("work failed:", err)
		}
	}()
	do(work)
}

func do(work *Work) {
	// 고의적으로 패닉을 발생시키는 예제
	panic("something went wrong with task: " + work.Task)
}

func main() {
	workChan := make(chan *Work, 10)

	go server(workChan)

	workChan <- &Work{Task: "Task 1"}
	workChan <- &Work{Task: "Task 2"}
	close(workChan)

	// 서버 고루틴이 완료될 때까지 잠시 대기
	select {}
}
