package main

import (
	"fmt"
	"time"
)

func main() {
	/*ch := getEmptyChannel()
	fmt.Println("start to select")

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("read from timer")
	case <-ch:
		fmt.Println("read from nil channel")
	}*/


	ch := make(chan int)

	for i := 0; i < 100; i++ {
		go func(i int) {
			for {
				req := <- ch
				fmt.Printf("worker[%d] got v from channel is [%d]\n", i, req)
			}
		}(i)
	}

	fmt.Println("start to write v into channel.")
	time.Sleep(2 * time.Second)

	//go func() {
	//	v := 0
	//	for {
	//
	//		ch <- v
	//		v++
	//		//time.Sleep(2 * time.Second)
	//	}
	//}()
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	//time.Sleep(1 * time.Second)
	ch <- 5
	ch <- 6
	ch <- 7
	ch <- 8
	//time.Sleep(1 * time.Second)
	ch <- 9
	ch <- 10
	ch <- 11
	ch <- 12
	//time.Sleep(1 * time.Second)

	time.Sleep(1 * time.Second)
}

func getEmptyChannel() <-chan struct{}{
	return nil
}
