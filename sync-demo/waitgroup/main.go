package main

import (
	"fmt"
	"sync"
	"time"
)

func add_num(a, b int, done func()) {
	defer done()
	c := a + b
	time.Sleep(1 * time.Second)
	fmt.Printf("%d + %d = %d\n", a, b, c)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go add_num(i, 1, wg.Done)
	}
	fmt.Printf("wait for 10 goroutine done.")
	wg.Wait()
	fmt.Printf("main done.")
}