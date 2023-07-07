package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	go func() {
		fmt.Printf("%d: ready call ctx.Done in one goroutine\n", time.Now().Second())
		<-ctx.Done()
		fmt.Printf("%d: complete steps in one goroutine\n", time.Now().Second())
	}()

	fmt.Printf("%d: will call cancelFunc after 3s sleep\n", time.Now().Second())
	time.Sleep(3 * time.Second)
	fmt.Printf("%d: do call cancelFunc\n", time.Now().Second())
	cancelFunc()

	time.Sleep(2 * time.Second)
	fmt.Printf("%d: main over after 2s sleep\n", time.Now().Second())
}
