package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Car struct {
	Body  string
	Tire  string
	Color string
}

var wg sync.WaitGroup
var startTime = time.Now()

func PaintCar(paintCh chan *Car) {
	for car := range paintCh {
		time.Sleep(time.Second)
		car.Color = "red"
		duration := time.Since(startTime)
		fmt.Printf("%.2f Complete Car: %s, %s, %s\n", duration.Seconds(), car.Body, car.Tire, car.Color)
	}
	wg.Done()
}

func InstallTire(tireCh chan *Car, paintCh chan *Car) {
	for car := range tireCh {
		time.Sleep(time.Second)
		car.Tire = "summer tire"
		paintCh <- car
	}
	close(paintCh)
	wg.Done()
}

func MakeBody(tireCh chan *Car) {
	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second)

	for {
		select {
		case <-tick:
			car := &Car{Body: "sports car"}
			tireCh <- car
		case <-after:
			close(tireCh)
			wg.Done()
			return
		}
	}
}

func PrintEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-ctx.Done():
			wg.Done()
			return
		}

	}
}

func main() {
	wg.Add(1)

	type contextKey string

	const numberKey contextKey = "number"

	ctx := context.WithValue(context.Background(), numberKey, 9)

	go square(ctx)

	wg.Wait()
}

func square(ctx context.Context) {
	if v := ctx.Value("number"); v != nil {
		n := v.(int)
		fmt.Println(n * n)
		wg.Done()
		return
	}
}
