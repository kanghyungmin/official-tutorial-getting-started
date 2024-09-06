package main

// var wg sync.WaitGroup

// func main() {
// 	rand.Seed(time.Now().UnixNano())

// 	wg.Add(2)

// 	fork := &sync.Mutex{}
// 	spoon := &sync.Mutex{}

// 	go diningProblem("A", fork, spoon, "포크", "수저")
// 	go diningProblem("B", spoon, fork, "수저", "포크")
// 	wg.Wait()
// }

// func diningProblem(name string, first, second *sync.Mutex, firstName, secondName string) {
// 	for i := 0; i < 100; i++ {
// 		fmt.Printf("%s는 식사 준비 중\n", name)

// 		first.Lock()
// 		fmt.Printf("%s는 %s를 들었습니다.\n", name, firstName)
// 		second.Lock()
// 		fmt.Printf("%s는 %s를 들었습니다.\n", name, secondName)

// 		fmt.Printf("%s가 식사 중\n", name)
// 		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

// 		second.Unlock()
// 		first.Unlock()

// 	}
// 	wg.Done()
// }

import (
	"fmt"
	"sync"
	"time"
)

type Job interface {
	Do()
}

type SquareJob struct {
	index int
}

func (j *SquareJob) Do() {
	fmt.Printf("%d 작업 시작\n", j.index)
	time.Sleep(1 * time.Second)
	fmt.Printf("%d 작업 완료 - 결과: %d\n", j.index, j.index*j.index)
}

func main() {
	var jobList [10]Job

	for i := 0; i < 10; i++ {
		jobList[i] = &SquareJob{i}
	}
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		// job := jobList[i]
		go func() {
			// job.Do()
			jobList[i].Do()
			wg.Done()
		}()
	}
	wg.Wait()
}
