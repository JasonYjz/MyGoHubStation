package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	//defer cancel()
	//go SlowOperation(ctx)
	//go func() {
	//	for {
	//		time.Sleep(300 * time.Millisecond)
	//		fmt.Println("goroutine:", runtime.NumGoroutine())
	//	}
	//}()
	//time.Sleep(4 * time.Second)

	fmt.Println("start main function.")
	ctx, cancelFunc := context.WithCancel(context.Background())
	go SlowOperation(ctx)

	fmt.Println("sleep 1 second.")
	time.Sleep(1 * time.Second)

	fmt.Println("do cancel func.")
	cancelFunc()

	time.Sleep(3)
}

func SlowOperation(ctx context.Context) {
	fmt.Println("start slow operation.")
	done := make(chan int, 1)
	go func() { // 模拟慢操作
		dur := time.Duration(5) * time.Second
		time.Sleep(dur)
		done <- 1
	}()

	select {
	case <-ctx.Done():
		fmt.Println("SlowOperation timeout:", ctx.Err())
	case <-done:
		fmt.Println("Complete work")
	}
}
