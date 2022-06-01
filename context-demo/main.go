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
	fmt.Println("new cancel context.")
	//ctx1, cancelFunc1 := context.WithTimeout(context.Background(), 2*time.Second)
	ctx1, cancelFunc1 := context.WithCancel(context.Background())
	fmt.Println("new 3s timeout context.")
	ctx2, cancelFunc2 := context.WithTimeout(ctx1, 3*time.Second)

	go SlowOperation(ctx2)

	defer cancelFunc1()
	defer cancelFunc2()

	fmt.Println("main start to sleep 5s")
	time.Sleep(5 * time.Second)
	//ctx, cancelFunc := context.WithCancel(context.Background())
	//go SlowOperation1(ctx)
	//go SlowOperation2(ctx)
	//
	//fmt.Println("sleep 1 second.")
	//time.Sleep(1 * time.Second)
	//
	//fmt.Println("do cancel func.")
	//cancelFunc()
	//
	//time.Sleep(3)
}

func SlowOperation1(ctx context.Context) {
	fmt.Println("start slow operation1.")
	done := make(chan int, 1)
	go func() { // 模拟慢操作
		dur := time.Duration(5) * time.Second
		time.Sleep(dur)
		done <- 1
	}()

	select {
	case <-ctx.Done():
		fmt.Println("SlowOperation1 timeout:", ctx.Err())
	case <-done:
		fmt.Println("Complete work 1 ")
	}
}

func SlowOperation2(ctx context.Context) {
	fmt.Println("start slow operation2.")
	done := make(chan int, 1)
	go func() { // 模拟慢操作
		dur := time.Duration(5) * time.Second
		time.Sleep(dur)
		done <- 1
	}()

	select {
	case <-ctx.Done():
		fmt.Println("SlowOperation2 timeout:", ctx.Err())
	case <-done:
		fmt.Println("Complete work 2")
	}
}

func SlowOperation(ctx context.Context) {
	fmt.Println("start to slow operation.")
	done := make(chan int, 1)
	go func() {
		time.Sleep(10 * time.Second)
		done <- 1
	}()

	select {
	case <-ctx.Done():
		fmt.Println("SlowOperation timeout", ctx.Err())
	case <-done:
		fmt.Println("Complete work")
	}
}
