package main

import (
	"fmt"
	"time"
)

func main() {
	ch := getEmptyChannel()
	fmt.Println("start to select")

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("read from timer")
	case <-ch:
		fmt.Println("read from nil channel")
	}
}

func getEmptyChannel() <-chan struct{}{
	return nil
}
